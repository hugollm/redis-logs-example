package main

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var redisClient redis.Conn
var logChan chan string

func log(msg string) {
	if logChan == nil {
		logChan = make(chan string)
	}
	logChan <- msg
}

func logHandler() {
	initRedisClient()
	for {
		msg := <-logChan
		addLogMessage(msg)
		printLatestLogs()
	}
}

func initRedisClient() {
	var err error
	redisClient, err = redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
}

func addLogMessage(msg string) {
	now := time.Now().Unix()
	logId := getNexLogId()
	msg = "[" + logId + "] " + msg
	_, err := redisClient.Do("ZADD", "log:entries", now, msg)
	if err != nil {
		panic(err)
	}
}

func getNexLogId() string {
	logId, err := redis.Int(redisClient.Do("INCR", "log:count"))
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(logId)
}

func printLatestLogs() {
	logs := getLatestLogs()
	print("\033[H\033[2J")
	for i := len(logs) - 1; i >= 0; i-- {
		println(logs[i])
	}
}

func getLatestLogs() []string {
	logs, err := redis.Strings(redisClient.Do("ZREVRANGE", "log:entries", 0, 9))
	if err != nil {
		panic(err)
	}
	return logs
}
