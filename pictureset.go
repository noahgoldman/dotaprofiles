package main

import (
	"fmt"
	"log"
	"path/filepath"
)

type PictureSet struct {
	id       int
	original string
	set      []string
}

func getPictureSet(id int) (*PictureSet, error) {
	original, err := Db.Cmd("lindex", id, "0").Str()
	if err != nil {
		return nil, fmt.Errorf("Failed to find a PictureSet with id=%d", id)
	}

	llen, err := Db.Cmd("llen", id).Int()
	if err != nil {
		return nil, err // TODO figure out a solution for this case
	}

	var set []string = nil
	// The number of elements when the set has been created is 6
	if llen == 6 {
		set, err = Db.Cmd("lrange", id, 1, 5).List()

		if err != nil {
			return nil, fmt.Errorf("The pictureset has an incorrect number of elements")
		}
	}

	return &PictureSet{id, original, set}, nil
}

func newPictureSet(file string) (*PictureSet, error) {
	id, err := Db.Cmd("incr", "next_id").Int()
	if err != nil {
		return nil, err
	}

	exists, err := Db.Cmd("exists", id).Bool()
	if err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("There is already data in id %d", id)
	}

	filename := fmt.Sprintf("%d_%s", id, file)

	llen, err := Db.Cmd("lpush", id, filename).Int()
	if err != nil {
		return nil, err
	} else if llen != 1 {
		panic("list length should not have changed...")
	}

	return &PictureSet{id, filename, nil}, nil
}

func (ps *PictureSet) AddSet() error {
	llen, err := Db.Cmd("llen", ps.id).Int()
	if err != nil {
		return err
	} else if llen < 1 {
		log.Fatal("list length error")
	}

	ps.set = make([]string, 5, 5)
	for i := 0; i < 5; i++ {
		ps.set[i] = fmt.Sprintf("%s_%d%s", GetFilenameWithoutExtension(ps.original),
			i+1, filepath.Ext(ps.original))
	}

	Db.Cmd("ltrim", ps.id, 0, 0)
	llen, err = Db.Cmd("rpush", ps.id, ps.set).Int()
	if err != nil {
		log.Fatal(err)
	} else if llen != 6 {
		log.Fatal("List length error while commiting")
	}

	return nil
}
