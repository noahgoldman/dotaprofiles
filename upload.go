package main

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"mime"
	"path/filepath"
)

const (
	BUCKET_NAME = "dotapics"
)

var bucket *s3.Bucket

func AWSInit() {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	s := s3.New(auth, aws.APNortheast)
	bucket = s.Bucket(BUCKET_NAME)
}

func Get_Mime(file string) (string, error) {
	ext := filepath.Ext(file)
	if ext == "" {
		return "", fmt.Errorf("Failed to get file extension of %s", file)
	}

	mime_t := mime.TypeByExtension(ext)
	if mime_t == "" {
		return "", fmt.Errorf("Failed to fine a mimetype for %s", file)
	}

	return mime_t, nil
}

func Get_New_Name(file string, new_name string) (string, error) {
	ext := filepath.Ext(file)
	if ext == "" {
		return "", fmt.Errorf("Failed to get file extension of %s", file)
	}

	return (new_name + ext), nil
}

func Upload_S3(file string, new_name string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Println("%#v", data)
	return err
}
