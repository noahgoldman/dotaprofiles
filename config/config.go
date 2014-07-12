package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Configuration struct {
	HttpPort      string
	AWS_AccessKey string
	AWS_SecretKey string
	DB            string
}

const (
	DEFAULT_FILE = "conf.json"
)

func GetDefaultFile() io.Reader {
	file, err := os.Open(DEFAULT_FILE)
	if err != nil {
		log.Fatal("Failed to find configuration file %s", DEFAULT_FILE)
	}

	return file
}

func LoadConfig(data io.Reader) *Configuration {
	decoder := json.NewDecoder(data)
	config := &Configuration{}

	err := decoder.Decode(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
