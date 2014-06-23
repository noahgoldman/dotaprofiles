package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
)

const (
	NETWORK    = "tcp"
	REDIS_ADDR = "localhost:6379"
)

var Db *redis.Client

func InitDB() {
	var err error
	Db, err = redis.Dial(NETWORK, REDIS_ADDR)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
