package study

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(_ context.Context, bucket string, key string, data multipart.File) error {

	fullPath := filepath.Join(bucket, key)

	// Ensure the parent directory exists
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	// Create the file
	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy file contents
	_, err = io.Copy(dst, data)
	return err
}
