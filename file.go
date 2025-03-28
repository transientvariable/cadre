package schema

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/transientvariable/support"

	json "github.com/json-iterator/go"
	gofs "io/fs"
)

// File godoc
// @Description Represents metadata for a file/object on a local/remote file system or storage service.
type File struct {
	Accessed   *time.Time `json:"accessed,omitempty" swaggerignore:"true"`
	Attributes []string   `json:"attributes,omitempty" swaggerignore:"true"`
	CID        string     `json:"cid,omitempty"`
	Ctime      *time.Time `json:"ctime,omitempty"`
	Created    *time.Time `json:"created,omitempty"`
	Directory  string     `json:"directory,omitempty"`
	Extension  string     `json:"extension,omitempty" swaggerignore:"true"`
	GID        string     `json:"gid,omitempty" swaggerignore:"true"`
	Group      string     `json:"group,omitempty" swaggerignore:"true"`
	Hash       *Hash      `json:"hash,omitempty"`
	Inode      string     `json:"inode,omitempty" swaggerignore:"true"`
	MimeType   string     `json:"mime_type,omitempty"`
	Mode       string     `json:"mode,omitempty" swaggerignore:"true"`
	Mtime      *time.Time `json:"mtime,omitempty"`
	Name       string     `json:"name,omitempty"`
	Owner      string     `json:"owner,omitempty" swaggerignore:"true"`
	Path       string     `json:"path,omitempty"`
	Size       int64      `json:"size,omitempty"`
	Type       string     `json:"type,omitempty"`
	UID        string     `json:"uid,omitempty" swaggerignore:"true"`
	URL        string     `json:"url,omitempty" swaggerignore:"true"`

	content []byte
} // @name File

// HashOf returns the value for the specified hash algorithm for the File.
//
// The zero-value will be returned if the File Hash is nil, the provided algorithm is empty, or does not match one of
// the algorithms for File.Hash.
func (f *File) HashOf(alg string) string {
	if f.Hash != nil {
		switch strings.ToLower(strings.TrimSpace(alg)) {
		case "adler", "adler32":
			return f.Hash.Adler32
		case "md5":
			return f.Hash.Md5
		case "sha1":
			return f.Hash.Sha1
		case "sha256":
			return f.Hash.Sha256
		case "sha512":
			return f.Hash.Sha512
		case "ssdeep":
			return f.Hash.Ssdeep
		}
	}
	return ""
}

// Content returns the File content. If the size of the File content > 0, then the returned byte slice will be
// a copy of the content, otherwise it will be the zero value.
func (f *File) Content() ([]byte, error) {
	if len(f.content) > 0 {
		c := make([]byte, len(f.content))
		_, err := io.Copy(bytes.NewBuffer(c), bytes.NewBuffer(f.content))
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// FileMode returns the os.FileMode for the File.
func (f *File) FileMode() gofs.FileMode {
	mode, err := strconv.Atoi(f.Mode)
	if err != nil {
		return os.ModeIrregular
	}
	return os.FileMode(mode)
}

// IsDir returns whether the File represents a regular file or directory.
func (f *File) IsDir() bool {
	return f.FileMode().IsDir()
}

// ToMap converts the File fields and their values to a map.
func (f *File) ToMap() (map[string]any, error) {
	var fm map[string]any
	if err := json.Unmarshal(support.ToJSON(f), &fm); err != nil {
		return nil, err
	}
	return fm, nil
}

// SetContent sets the File content. If the size of the provided content > 0, then File content will be set to  a copy.
func (f *File) SetContent(c []byte) error {
	if len(c) > 0 {
		f.content = make([]byte, len(c))
		s, err := io.Copy(bytes.NewBuffer(f.content), bytes.NewReader(c))
		if err != nil {
			return err
		}
		f.Size = s
	}
	return nil
}
