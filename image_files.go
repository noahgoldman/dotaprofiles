package main

import (
	"path"
	"os"
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
func GetImageFile(name string) (*os.File, error) {
	return os.Open(GetImgPath(name))
}
