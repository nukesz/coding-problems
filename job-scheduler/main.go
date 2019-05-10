package main

import (
	"log"
	"time"
)

func schedule(f func(), n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
	f()
}

// Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
func main() {
	schedule(printHello, 1)
}

func printHello() {
	log.Println("Hello Scheduler")
}
