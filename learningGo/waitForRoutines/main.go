package main

import (
	"fmt"
	"sync"
    "time"
)

func doStuff() {
    time.Sleep(500 * time.Millisecond)
}

func main() {
    fmt.Println("Starting...")
    var wg sync.WaitGroup
    wg.Add(3)
    for i := 0; i < 3; i++ {
        go func(i int) {
            defer wg.Done()
            fmt.Printf("Goroutine #%d\n", i)
            doStuff()
        }(i)
    }
    wg.Wait()
    fmt.Println("All done")
}
