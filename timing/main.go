package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()

	log.Println("Current time is: ", now.Format("2006-01-02 15:04:05"))

	time.AfterFunc(2*time.Second, func() {
		log.Println("2 seconds later than above...")
	})
	time.Sleep(3 * time.Second)

	total := time.Since(now)
	log.Println("Total time in seconds: ", total.Seconds())
}
