package config

import (
	"bytes"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	test_data := `{
		"HttpPort": ":8080",
		"AWS_AccessKey": "accesscode",
		"AWS_SecretKey": "password",
		"DB": "localhost:5000"
	}`

	buffer := bytes.NewBufferString(test_data)
	config := LoadConfig(buffer)

	if config.HttpPort != ":8080" ||
		config.AWS_SecretKey != "password" {
		t.Error("Failed to parse a config correctly")
	}
}
