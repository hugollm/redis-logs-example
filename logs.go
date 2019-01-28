package main

import (
	"strconv"
	"time"
)

var chlog chan string = make(chan string)

func log(msg string) {
	chlog <- msg
}

func logHandler() {
	for {
		msg := <-chlog
		addLogMessage(msg)
		printLatestLogs()
	}
}

func addLogMessage(msg string) {
	id := incr("log:count")
	now := int(time.Now().UnixNano())
	msg = "[" + strconv.Itoa(id) + "] " + msg
	zadd("log:entries", now, msg)
}

func printLatestLogs() {
	logs := zrevrange("log:entries", 0, 99)
	print("\033[H\033[2J")
	for i := len(logs) - 1; i >= 0; i-- {
		println(logs[i])
	}
}
