package storage

import (
	"fmt"

	"github.com/transientvariable/schema-go"

	"github.com/minio/sha256-simd"
)

func fileID(file *schema.File) string {
	if file == nil || file.Path == "" {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(file.Path)))
}
