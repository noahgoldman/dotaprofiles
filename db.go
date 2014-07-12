package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
)

const (
	NETWORK = "tcp"
)

var Db *redis.Client

func InitDB() {
	var err error
	Db, err = redis.Dial(NETWORK, Config.DB)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
