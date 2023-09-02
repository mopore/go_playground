package mqtt

import (
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

const (
    stopic = "jniHome/services/telegramBot/aliveTick"
    ptopic = "goTest"
    brokerAddress = "tcp://192.168.199.119:1883"
    clientId = "goTestClient"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    log.Printf("TOPIC: %s\n", msg.Topic())
    log.Printf("MSG: %s\n", msg.Payload())
}

func PerformMqttRun() {
    log.Println("Starting MQTT run")
    opts := mqtt.NewClientOptions().AddBroker(brokerAddress)
    opts.SetClientID(clientId)
    opts.SetKeepAlive(2 * time.Second)
    opts.SetDefaultPublishHandler(f)
    opts.SetPingTimeout(1 * time.Second)

    c := mqtt.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    if token := c.Subscribe(stopic, 0, nil); token.Wait() && token.Error() != nil{
        panic(token.Error())
    }

    for i := 0; i < 5; i++ {
        text := fmt.Sprintf("This is msg #%d!", i)
        token := c.Publish(ptopic, 0, false, text)
        token.Wait()
    }

    time.Sleep(6 * time.Second)

    if token := c.Unsubscribe(stopic); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    c.Disconnect(250)

    time.Sleep(1 * time.Second)
    log.Println("Finishing MQTT run without errors")
}
