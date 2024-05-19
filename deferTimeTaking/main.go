package main

import (
	"log"
	"time"
)

func main() {
    start := time.Now()
    defer func() {
        log.Println("Time taken: ", time.Since(start))
    }() // It is important to call!

    log.Println("Starting to time")
    time.Sleep(2 * time.Second)
}
