package main

import (
	"github.com/fzzy/radix/redis"
	"fmt"
)

const (
	NETWORK = "tcp"
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
