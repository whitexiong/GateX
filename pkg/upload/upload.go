package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SaveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return "", err
	}

	destPath := filepath.Join(destDir, file.Filename)
	out, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	destPath = strings.ReplaceAll(destPath, "\\", "/")
	return destPath, nil
}
