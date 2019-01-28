package main

import (
	"time"
)

func main() {
	go logHandler()
	for {
		log("lorem ipsum")
		time.Sleep(100 * time.Millisecond)
	}
}
