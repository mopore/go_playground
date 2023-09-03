package main

import (
	"fmt"
	"sync"
    "time"
)

func main() {
    fmt.Println("Starting...")
    var wg sync.WaitGroup
    wg.Add(3)
    for i := 0; i < 3; i++ {
        go func(i int) {
            defer wg.Done()
            fmt.Printf("Goroutine #%d\n", i)
            time.Sleep(500 * time.Millisecond)
        }(i)
    }
    wg.Wait()
    fmt.Println("All done")
}
