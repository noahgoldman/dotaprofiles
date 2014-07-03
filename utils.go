package main

import (
	"image"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func GetRect(r *http.Request) (*image.Rectangle, error) {
	vars := [...]string{"x1", "y1", "x2", "y2"}
	var out_vars [4]int

	for key, value := range vars {
		value, err := strconv.ParseFloat(r.FormValue(value), 64)
		if err != nil {
			return nil, err
		}

		out_vars[key] = int(value)
	}

	return &image.Rectangle{image.Pt(out_vars[0], out_vars[1]),
		image.Pt(out_vars[2], out_vars[3])}, nil
}

func GetFilenameWithoutExtension(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(path, filepath.Ext(base))
}
