package main

import (
	"combo/processor"
	"context"
	"fmt"
    "time"
)

func gatherAndProcess(ctx context.Context, inA, inB int) (string, error) {
    ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
    defer cancel()
    p := processor.Processor{
        OutA: make(chan int, 1),
        OutB: make(chan int, 1),
        InC: make(chan int, 1),
        OutC: make(chan string, 1),
        Errs: make(chan error, 2),
    }
    p.Launch(ctx, inA, inB)
    inputC, err := p.WaitForAB(ctx)
    if err != nil {
        return "", err
    }
    p.InC <- inputC
    out, err := p.WaitForC(ctx)
    return out, err
}


func main() {
    fmt.Println("Hello, World!")
    inA := 1
    inB := 200
    ctx := context.Background()
    out, err := gatherAndProcess(ctx, inA, inB)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out)
    fmt.Println("Goodbye, World!")
}
