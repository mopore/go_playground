package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const routineCount = 3

func doStuff() int {
    // Create a random number between 100 and 500
    r := rand.Intn(400) + 100
    time.Sleep(time.Duration(r) * time.Millisecond)
    return r
}

func main() {
    fmt.Printf("Preparing to run %d goroutines.\n", routineCount)
    var wg sync.WaitGroup
    results := make(chan int, routineCount)
    wg.Add(routineCount)
    for i := 0; i < routineCount; i++ {
        go func(i int) {
            defer wg.Done()
            fmt.Printf("Goroutine #%d\n", i)
            r := doStuff()
            results <- r
        }(i)
    }
    wg.Wait()
    close(results)
    fmt.Println("All routines completed.")
    for r := range results {
        fmt.Printf("Got result: %d\n", r)
    }
}
