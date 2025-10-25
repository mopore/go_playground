package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const rCount = 6

func doStuff() int {
    r := rand.Intn(500-100) + 100  // Create a random number between 100 and 500
    time.Sleep(time.Duration(r) * time.Millisecond)
    return r
}

func main() {
    fmt.Printf("Preparing to run %d goroutines.\n", rCount)

    var wg sync.WaitGroup
    results := make(chan int, rCount)

    for i := range rCount {
        // Avoid old syntax with WaitGroup.Add and Waitgroup.Done
        wg.Go(func() {
            fmt.Printf("Goroutine #%d\n", i)
            r := doStuff()
            results <- r
        })
    }
    wg.Wait()
    close(results)

    fmt.Println("All routines completed.")
    for r := range results {
        fmt.Printf("Got result: %d\n", r)
    }
}
