package main

import (
	"github.com/stvp/tempredis"
	"github.com/fzzy/radix/redis"
	"testing"
	"fmt"
)

const (
	TESTING_PORT = "11001"
)

func setupDB() (*tempredis.Server) {
	server, err := tempredis.Start(
		tempredis.Config{
			"port": TESTING_PORT,
			"databases": "1",
		},
	)

	if err != nil {
		panic(err)
	}
	
	Db, err = redis.Dial(NETWORK, ":" + TESTING_PORT)

	return server
}

func TestNewPictureSet(t *testing.T) {
	server := setupDB()
	defer server.Term()	
	defer Db.Close()

	file := "testfile.jpg"

	ps, err := newPictureSet(file)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	if ps.id != 1 {
		fmt.Errorf("The id of a new PictureSet should be 1 not %d", ps.id)
	}	

	str, _ := Db.Cmd("lindex", 1, 0).Str()
	if str != file {
		fmt.Errorf("The file name should be %s not %s", file, str)
	}	
}
