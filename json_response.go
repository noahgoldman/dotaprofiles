package main

import (
	"encoding/json"
)

type JSONResponse map[string]interface{}

func (r JSONResponse) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)
}
