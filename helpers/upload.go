package helpers

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader, path string) (err error) {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
