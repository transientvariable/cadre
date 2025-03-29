package car

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/transientvariable/schema-go"
	"github.com/transientvariable/support-go"

	json "github.com/json-iterator/go"
)

const (
	MetadataFileName = "metadata.json"
	EntriesFileName  = "entries.csv"
	EntriesCSVFields = "name,path,size,sha256,mtime"
)

type Metadata struct {
	Entries   int    `json:"entries"`
	Index     int    `json:"page"`
	Namespace string `json:"namespace"`
	Size      int64  `json:"size"`
}

type Manifest struct {
	entries     []*schema.File
	entriesPath string
	graphsplit  GraphsplitManifest
	metadata    *Metadata
	mutex       sync.RWMutex
	path        string
}

func NewManifest(namespace string, index uint) *Manifest {
	return &Manifest{
		metadata: &Metadata{
			Namespace: namespace,
			Index:     int(index),
		},
	}
}

func (m *Manifest) Add(entries ...*schema.File) {
	for _, entry := range entries {
		m.entries = append(m.entries, entry)
		m.metadata.Size += entry.Size
		m.metadata.Entries += 1
	}
}

func (m *Manifest) Graphsplit() GraphsplitManifest {
	return m.graphsplit
}

func (m *Manifest) Count() int {
	return m.metadata.Entries
}

func (m *Manifest) EntryNames() []string {
	var names []string
	for _, e := range m.entries {
		names = append(names, e.Name)
	}
	return names
}

func (m *Manifest) Id() string {
	if m.Index() < 10 {
		return fmt.Sprintf("0%d", m.Index())
	}
	return strconv.Itoa(m.Index())
}

func (m *Manifest) Index() int {
	return m.metadata.Index
}

func (m *Manifest) Namespace() string {
	return m.metadata.Namespace
}

func (m *Manifest) Path() string {
	return m.path
}

func (m *Manifest) ReadAllEntries() ([]*schema.File, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.entries) > 0 {
		m.entries = m.entries[:]
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	entries, err := m.ReadEntries(ctx)
	if err != nil {
		return nil, err
	}

	for entry := range entries {
		m.Add(entry)
	}
	return m.entries, nil
}

func (m *Manifest) ReadEntries(ctx context.Context) (<-chan *schema.File, error) {
	entries := make(chan *schema.File)
	go func() {
		defer close(entries)

		f, err := os.Open(m.entriesPath)
		if err != nil {
			fmt.Println(fmt.Errorf("manifest: %w", err))
			return
		}
		defer func(f *os.File) {
			if err := f.Close(); err != nil {
				fmt.Println(fmt.Errorf("manifest: %w", err))
			}
		}(f)

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		if scanner.Scan() {
			_ = scanner.Text() // skip header row
		}

		for scanner.Scan() {
			attrs := strings.Split(scanner.Text(), ",")

			size, err := strconv.ParseInt(attrs[2], 10, 64)
			if err != nil {
				fmt.Println(fmt.Errorf("manifest: %w", err))
				return
			}

			mtime, err := time.Parse(time.RFC3339Nano, attrs[4])
			if err != nil {
				fmt.Println(fmt.Errorf("manifest: %w", err))
				return
			}

			select {
			case entries <- &schema.File{
				Name:  attrs[0],
				Path:  attrs[1],
				Size:  size,
				Hash:  &schema.Hash{Sha256: attrs[3]},
				Mtime: &mtime,
			}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return entries, nil
}

func (m *Manifest) Size() int64 {
	return m.metadata.Size
}

func (m *Manifest) WriteTo(dst string) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	dir := dir(dst, m.Index())
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filepath.Join(dir, MetadataFileName), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println(fmt.Errorf("manifest: %w", err))
		}
	}(file)

	if _, err = file.WriteString(m.String()); err != nil {
		return err
	}
	return m.writeEntriesTo(dir)
}

func (m *Manifest) String() string {
	pm := make(map[string]any)
	pm["metadata"] = m.metadata

	if len(m.graphsplit.Entries) > 0 {
		pm["graphsplit"] = m.graphsplit
	}
	return string(support.ToJSONFormatted(pm))
}

func (m *Manifest) copyEntry(e *schema.File) (*schema.File, error) {
	c, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	var entry *schema.File
	if err := json.NewDecoder(bytes.NewReader(c)).Decode(&entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (m *Manifest) writeEntriesTo(dst string) error {
	file, err := os.OpenFile(filepath.Join(dst, EntriesFileName), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	if _, err = file.WriteString(EntriesCSVFields + "\n"); err != nil {
		return err
	}

	sort.Slice(m.entries, func(i int, j int) bool { return m.entries[i].Mtime.Before(*m.entries[j].Mtime) })

	for _, e := range m.entries {
		fields := []string{
			e.Name,
			e.Path,
			strconv.FormatInt(e.Size, 10),
		}

		if e.Hash != nil {
			fields = append(fields, e.HashOf("sha256"))
		} else {
			fields = append(fields, "")
		}

		if !e.Mtime.IsZero() {
			fields = append(fields, e.Mtime.Format(time.RFC3339Nano))
		} else {
			fields = append(fields, "")
		}

		if _, err = file.WriteString(strings.Join(fields, ",") + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func ReadWithIndex(src string, index int) (*Manifest, error) {
	return Read(dir(src, index))
}

func Read(src string) (*Manifest, error) {
	if _, err := os.Stat(src); err != nil {
		return nil, err
	}

	graphsplit, err := NewGraphsplitManifest(filepath.Join(src, GraphsplitManifestFileName))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	metadata, err := readMetadata(src)
	if err != nil {
		return nil, err
	}
	return &Manifest{
		entriesPath: filepath.Join(src, EntriesFileName),
		graphsplit:  graphsplit,
		metadata:    metadata,
		path:        src,
	}, nil
}

func readMetadata(dir string) (*Metadata, error) {
	b, err := os.ReadFile(filepath.Join(dir, MetadataFileName))
	if err != nil {
		return nil, err
	}

	var mm map[string]any
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&mm); err != nil {
		return nil, err
	}

	var metadata *Metadata
	if m, ok := mm["metadata"]; ok {
		if err := json.NewDecoder(bytes.NewReader(support.ToJSON(m))).Decode(&metadata); err != nil {
			return nil, err
		}
	} else {
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&metadata); err != nil {
			return nil, err
		}
	}
	return metadata, nil
}

func dir(path string, index int) string {
	var dir string
	if index < 10 {
		dir = filepath.Join(path, fmt.Sprintf("0%d", index))
	} else {
		dir = filepath.Join(path, strconv.Itoa(index))
	}
	return dir
}
