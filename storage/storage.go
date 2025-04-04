package storage

import (
	"fmt"

	"github.com/transientvariable/cadre"

	"github.com/minio/sha256-simd"
)

func fileID(file *cadre.File) string {
	if file == nil || file.Path == "" {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(file.Path)))
}
