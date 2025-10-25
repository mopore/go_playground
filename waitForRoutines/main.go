package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const rCount = 6

func doStuff() int {
    // Create a random number between 100 and 500
    r := rand.Intn(500-100) + 100
    time.Sleep(time.Duration(r) * time.Millisecond)
    return r
}

func main() {
    fmt.Printf("Preparing to run %d goroutines.\n", rCount)
    var wg sync.WaitGroup
    results := make(chan int, rCount)

    // // old syntax
    // wg.Add(rCount)
    // for i := 0; i < rCount; i++ {
    //     go func(i int) {
    //         defer wg.Done()
    //         fmt.Printf("Goroutine #%d\n", i)
    //         r := doStuff()
    //         results <- r
    //     }(i)
    // }

    // NEW syntax
    for i := range rCount {
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
