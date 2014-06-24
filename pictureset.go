package main

import (
	"log"
	"github.com/fzzy/radix/redis"
	"fmt"
)

type PictureSet struct {
	id       int
	original string
	set      []string
}

func getPictureSet(id int) (*PictureSet, error){

	original, err := Db.Cmd("lindex", id, "0").Str()
	if err != nil {
		return nil, fmt.Errorf("Failed to find a PictureSet with id=%d", id)
	}

	llen, err := Db.Cmd("llen", id).Int()
	if err != nil {
		return nil, err
	}

	var set []string = nil
	// The number of elements when the set has been created is 6
	if llen == 6 {
		set, err = Db.Cmd("lrange", id, 1, 5).List()

		if err != nil {
			log.Fatal(err)
		}
	}

	return &PictureSet{id, original, set}, nil
}

func newPictureSet(file string) (*PictureSet, error) {
	r := Db.Cmd("get", "next_id")
	id, err := r.Int()
	if r.Type == redis.NilReply {
		r = Db.Cmd("set", "next_id", 0)
		id = 0
	} else if r.Type == redis.ErrorReply {
		log.Fatal("database conn failed")
	}

	exists, err := Db.Cmd("exists", id).Bool()
	if exists || err != nil {
		log.Fatal(err)
	}

	llen, err := Db.Cmd("lpush", id, file).Int()
	if err != nil {
		log.Fatal(err)
	} else if llen != 1 {
		panic("list length should not have changed...")
	}

	return &PictureSet{id, file, nil}, nil
}

func (ps *PictureSet) addSet(set []string) error {
	llen, err := Db.Cmd("llen", ps.id).Int()
	if err != nil {
		return err
	} else if llen < 1 {
		log.Fatal("list length error")
	}

	Db.Cmd("ltrim", 0, 0)
	llen, err = Db.Cmd("rpush", set).Int()
	if err != nil {
		log.Fatal(err)
	} else if llen != 6 {
		log.Fatal("List length error while commiting")
	}

	return nil
}
