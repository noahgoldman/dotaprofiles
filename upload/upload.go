package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"mime"
	"path/filepath"
)

const (
	BUCKET_NAME = "dotaprofiles"
)

var PicsBucket *s3.Bucket

func AWSInit() error {
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	s := s3.New(auth, aws.USEast)
	PicsBucket = s.Bucket(BUCKET_NAME)
	return nil
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

func Upload_S3(file io.Reader, new_name string) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	mime_t, err := Get_Mime(new_name)
	if err != nil {
		return err
	}

	err = PicsBucket.Put(new_name, data, mime_t, s3.PublicRead)
	return err
}
