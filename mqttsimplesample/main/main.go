package main

import (
    "log"
    "mqttsimplesample/mqtt"
)

func main() {
    log.Println("Entry of main...")
    mqtt.PerformMqttRun()
    log.Println("Exit of main.")
}

