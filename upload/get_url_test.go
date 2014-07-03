package upload

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetURL(t *testing.T) {
	filename := UploadTestFile(t)

	url := GetURL(filename)

	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(body) != TEST_FILE_DATA {
		t.Errorf("HTTP response %s is not %s", string(body), TEST_FILE_DATA)
	}
}
