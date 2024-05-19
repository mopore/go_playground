package main

// countTo returns a channel that will send the numbers from 1 to max, and a
// function that can be called to cancel the count.
//
// This can also be viewed as a generator function.
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
            case ch <- i: // Just send the number into the channel
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
