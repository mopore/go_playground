package main

import (
    "log"
    "mqttsimplesample/mqtt"
    "sync"
)

func main() {
    log.Println("Main start.")

    var wg sync.WaitGroup
    wg.Add(1)
    go mqtt.RunMqtt(&wg)

    log.Println("Main waits now...")
    wg.Wait()

    log.Println("Main exit.")
}

