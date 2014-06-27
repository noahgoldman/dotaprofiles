package main

import (
	"io"
	"os"
	"path"
)

const (
	IMAGES_DIR = "img_store"
)

func GetImgPath(name string) string {
	return path.Join(IMAGES_DIR, name)
}

func CreateImageFile(name string) (*os.File, error) {
	return os.Create(GetImgPath(name))
}

// Load a file from the directory
func GetImageFile(name string, download_file func(string) (io.Reader, error)) (io.Reader, error) {
	file, err := os.Open(GetImgPath(name))
	if err == nil {
		return file, nil
	}

	return download_file(name)
}
