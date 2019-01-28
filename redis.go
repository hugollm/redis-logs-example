package main

import (
	"github.com/gomodule/redigo/redis"
)

var rclient redis.Conn

func init() {
	rclient = connect()
}

func connect() redis.Conn {
	client, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return client
}

func incr(key string) int {
	count, err := redis.Int(rclient.Do("INCR", key))
	if err != nil {
		panic(err)
	}
	return count
}

func zadd(key string, score int, value string) {
	_, err := rclient.Do("ZADD", key, score, value)
	if err != nil {
		panic(err)
	}
}

func zrevrange(key string, start int, stop int) []string {
	entries, err := redis.Strings(rclient.Do("ZREVRANGE", key, start, stop))
	if err != nil {
		panic(err)
	}
	return entries
}
