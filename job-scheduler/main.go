package main

import (
	"log"
	"time"
)

// Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
func schedule(f func(), n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
	f()
}

func main() {
	schedule(printHello, 2000)
}

func printHello() {
	log.Println("Hello Scheduler")
}
