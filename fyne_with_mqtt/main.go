package main

import (
    "log"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

const (
    brokerAddress = "tcp://192.168.199.119:1883"
)

func main() {
    log.Println("Starting the application...")
    a := app.New()
    w := a.NewWindow("Fyne MQTT Test")
    w.SetContent(widget.NewLabel("Hello Viewer!"))
    w.ShowAndRun()
}
