package main

import (
	"testing"
)

func TestGetFilenameWithoutExtension(t *testing.T) {
	testpath := "/home/local/test.png"
	res := GetFilenameWithoutExtension(testpath)
	if res != "test" {
		t.Errorf("The name should be %s not %s", "test", res)
	}

	testpath = "otherfile.txt"
	res = GetFilenameWithoutExtension(testpath)
	if res != "otherfile" {
		t.Errorf("The name should be %s not %s", "otherfile", res)
	}

	testpath = "/testdir/otherfile"
	res = GetFilenameWithoutExtension(testpath)
	if res != "otherfile" {
		t.Errorf("The name should be %s not %s", "otherfile", res)
	}
}
