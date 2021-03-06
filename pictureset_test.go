package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/stvp/tempredis"
	"testing"
)

const (
	TESTING_PORT = "11001"
)

func setupDB(t *testing.T) *tempredis.Server {
	server, err := tempredis.Start(
		tempredis.Config{
			"port":      TESTING_PORT,
			"databases": "1",
		},
	)

	if err != nil {
		t.Skip("Failed to connect to redis server, so skipping")
	}

	Db, err = redis.Dial(NETWORK, ":"+TESTING_PORT)

	return server
}

func TestNewPictureSet(t *testing.T) {
	server := setupDB(t)
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

func TestGetPictureSet(t *testing.T) {
	server := setupDB(t)
	defer server.Term()
	defer Db.Close()

	id := 1
	file := "test.jpg"

	llen, _ := Db.Cmd("lpush", id, file).Int()
	if llen != 1 {
		fmt.Errorf("Somehow the database operation didnt work, len should be 1 not %d", llen)
	}

	ps, _ := getPictureSet(id)
	if ps.original != file {
		fmt.Errorf("Get picture original should be %s not %s", file, ps.original)
	}
}
