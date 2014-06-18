package main

import (
	"log"
)

type PictureSet struct  {
	id int
	original string
	set []string
}

func getPictureSet(id int) *PictureSet {

	original, err := Db.Cmd("lindex", id, "0").Str()

	if err != nil {
		log.Fatal(err)
	}

	llen, err := Db.Cmd("llen", id).Int()
	if err != nil {
		log.Fatal(err)
	}

	var set []string = nil
	// The number of elements when the set has been created is 6
	if llen == 6 {
		set, err = Db.Cmd("lrange", id, 1, 5).List()
		
		if err != nil {
			log.Fatal(err)
		}
	}
	
	return &PictureSet{id, original, set}
}

func newPictureSet(file string) *PictureSet {
	id, err := Db.Cmd("get", "next_id").Int()
	if err != nil {
		log.Fatal(err)
	}

	exists, err := Db.Cmd("exists", id).Bool()
	if exists || err != nil {
		log.Fatal(err)
	}

	llen, err := Db.Cmd("lpush", file).Int()
	if llen != 1 {
		panic("list length should not have changed...")
	}
}

func (ps *PictureSet) makeSet(set []string) error {
	llen, err := Db.Cmd("llen", ps.id).Int()
	if err != nil {
		return err
	} else if llen < 1 {
		log.Fatal("list length error")
	}

	Db.Cmd("ltrim", 0, 0)
	llen, err = Db.Cmd("rpush", set)
	if err != nil {
		log.Fatal(err)
	} else if llen != 6 {
		log.Fatal("List length error while commiting")
	}
}
