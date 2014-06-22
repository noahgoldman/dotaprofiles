package main

import (
	"testing"
)

func TestGetMime(t *testing.T) {
	mime_t, err := Get_Mime("/local/testpic.png")
	if mime_t != "image/png" {
		if err == nil {
			t.Error("Error was not thrown for PNG")
		}
		t.Errorf("PNG mime type is not %s", mime_t)
	}

	mime_t, err = Get_Mime("testpic2.jpg")
	if mime_t != "image/jpeg" {
		if err == nil {
			t.Error("Error was not thrown for JPG")
		}
		t.Errorf("JPG mime type is not %s", mime_t)
	}
}

func TestGetMimeFail(t *testing.T) {
	_, err := Get_Mime("folder/testfile")
	if err == nil {
		t.Error("Should've thrown an error for no extension in Get_Mime")
	}

	_, err = Get_Mime("folder/testfile.fakeextensionfortesting")
	if err == nil {
		t.Error("Should've thrown an error for not MIME type in Get_Mime")
	}
}

func TestGetNewName(t *testing.T) {
	name, err := Get_New_Name("folder/testfile.jpg", "othertest")
	if name != "othertest.jpg" {
		t.Errorf("New name was not correct, it was %s instead of othertest.jpg", name)
	}

	name, err = Get_New_Name("folder/testfile", "othertest")
	if err == nil {
		t.Error("Get_New_Name should've failed")
	}
}
