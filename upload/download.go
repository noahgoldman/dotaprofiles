package upload

func Download_S3(file string) ([]byte, error) {
	return PicsBucket.Get(file)
}
