package main

import (
	"io"
	"os"
	"path"
	"bytes"
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
func GetImageFile(name string, download_file func(string) ([]byte, error)) (io.Reader, error) {
	file, err := os.Open(GetImgPath(name))
	if err == nil {
		return file, nil
	}

	data, err := download_file(name)
	if err != nil {
		return nil, err
	}	

	return bytes.NewBuffer(data), nil
}
