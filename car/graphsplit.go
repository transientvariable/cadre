package car

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/transientvariable/anchor"
	"github.com/transientvariable/cadre"

	"github.com/ipfs/go-cid"

	json "github.com/json-iterator/go"
	pool "github.com/libp2p/go-buffer-pool"
)

const (
	tokenBufferSize = anchor.MiB
	tokenSizeMax    = 10 * anchor.MiB

	GraphsplitManifestFileName = "manifest.csv"
)

var tokenBufferPool pool.BufferPool

type GraphsplitManifestEntry struct {
	FileName    string `json:"file_name"`
	PayloadCID  string `json:"payload_cid"`
	PayloadHash string `json:"payload_hash"`
	PayloadSize int64  `json:"payload_size"`
	PieceCID    string `json:"piece_cid"`
	PieceHash   string `json:"piece_hash"`
	PieceSize   int64  `json:"piece_size"`
}

func (e GraphsplitManifestEntry) ToMap() (map[string]any, error) {
	var m map[string]any
	if err := json.NewDecoder(strings.NewReader(e.String())).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}

func (e GraphsplitManifestEntry) String() string {
	return string(anchor.ToJSONFormatted(e))
}

type GraphsplitManifest struct {
	File    cadre.File
	Entries []GraphsplitManifestEntry
}

func NewGraphsplitManifest(path string) (GraphsplitManifest, error) {
	path = strings.TrimSpace(path)

	mp, err := os.Stat(path)
	if err != nil {
		return GraphsplitManifest{}, err
	}

	manifestFile := cadre.File{Path: path}
	if mp.IsDir() {
		manifestFile.Directory = path
		manifestFile.Name = GraphsplitManifestFileName
		manifestFile.Path = filepath.Join(path, GraphsplitManifestFileName)
	} else {
		manifestFile.Directory = filepath.Dir(path)
		manifestFile.Name = filepath.Base(path)
	}

	mf, err := os.Open(manifestFile.Path)
	if err != nil {
		return GraphsplitManifest{}, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Println(fmt.Errorf("car_manifest: %w", err))
		}
	}(mf)

	entries, err := readEntries(mf)
	if err != nil {
		return GraphsplitManifest{}, err
	}

	return GraphsplitManifest{
		File:    manifestFile,
		Entries: entries,
	}, nil
}

func readEntries(file *os.File) ([]GraphsplitManifestEntry, error) {
	var entries []GraphsplitManifestEntry

	scanner := bufio.NewScanner(file)
	buffer := tokenBufferPool.Get(tokenBufferSize)
	defer tokenBufferPool.Put(buffer)
	scanner.Buffer(buffer, tokenSizeMax)

	if scanner.Scan() {
		_ = scanner.Text() // skip header row
	}

	for scanner.Scan() {
		attrs := strings.Split(scanner.Text(), ",")

		payloadCID, err := cid.Parse(attrs[0])
		if err != nil {
			return nil, err
		}

		pieceCID, err := cid.Parse(attrs[2])
		if err != nil {
			return nil, err
		}

		payloadSize, err := strconv.ParseInt(attrs[3], 10, 64)
		if err != nil {
			return nil, err
		}

		pieceSize, err := strconv.ParseInt(attrs[4], 10, 64)
		if err != nil {
			return nil, err
		}

		entries = append(entries, GraphsplitManifestEntry{
			FileName:    attrs[1],
			PayloadCID:  attrs[0],
			PayloadHash: payloadCID.Hash().HexString(),
			PayloadSize: payloadSize,
			PieceCID:    attrs[2],
			PieceHash:   pieceCID.Hash().HexString(),
			PieceSize:   pieceSize,
		})
	}
	return entries, nil
}

func (m GraphsplitManifest) String() string {
	return string(anchor.ToJSONFormatted(m))
}
