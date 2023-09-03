package main

import (
	"errors"
	"fmt"
	"time"
    "log"
)

type Limiter struct {
    ch chan struct{}
}

func NewLimiter(limit int) *Limiter {
    ch := make(chan struct{}, limit)
    for i := 0; i < limit; i++ {
        ch <- struct{}{}
    }
    return &Limiter{
        ch: ch,
    }
}

func (pg *Limiter) Process(i int, f func(int)) error {
    select {
    case <- pg.ch:
        f(i)
        pg.ch <- struct{}{}
        return nil
    default:
        return errors.New("no more capacity")
    }
}


func main() {
    limiter := NewLimiter(4)
    for i := 0; i < 10; i++ {
        go func(i int) {
            err := limiter.Process(i, func(i int) {
                fmt.Printf("Starting...%d\n", i)
                time.Sleep(1 * time.Millisecond)
                fmt.Printf("Finishing...%d\n", i)
            })
            if err != nil {
                log.Fatal(err)
            }
        }(i)

        // Delay next go routine
        time.Sleep(1 * time.Millisecond)
    }
    time.Sleep(1 * time.Second)
    fmt.Println("Done")
}
