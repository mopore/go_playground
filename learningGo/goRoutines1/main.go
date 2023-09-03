package main

func countTo(max int) (<-chan int, func()) {
    ch := make(chan int)
    done := make(chan struct{})
    cancel := func () {
        close(done)
    }
    go func() {
        for i := 1; i <= max; i++ {
            select {
            case <-done:
                return
            case ch <- i:
            // nothing
            }
        }
        close(ch)
    }()
    return ch, cancel
}

func main() {
    ch, cancel := countTo(10)
    for i := range ch {
        if i > 5 {
            break
        }
        println(i)
    }
    cancel()
}
