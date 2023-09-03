package mqtt

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

const (
    aliveTopic = "goTest/services/mqttBridge/aliveTick"
    commandTopic = "goTest/services/mqttBridge/command"
    testTopic = "goTest/services/mqttBridge/test"
    brokerAddress = "tcp://192.168.199.119:1883"
    clientId = "goTestClient"
    aliveFrequency = 10 * time.Second
)

func publish5TestMessages(c mqtt.Client) {
    log.Printf("Publishing 5 test messages on topic %s...\n", testTopic)
    for i := 0; i < 5; i++ {
        text := fmt.Sprintf("This is msg #%d!", i+1)
        token := c.Publish(testTopic, 0, false, text)
        token.Wait()
    }
}

func RunMqtt(wg *sync.WaitGroup) {
    log.Println("Preparing MQTT client...")
    cdone := make(chan struct{})

    opts := mqtt.NewClientOptions().AddBroker(brokerAddress)
    opts.SetClientID(clientId)
    opts.SetKeepAlive(2 * time.Second)
    opts.SetPingTimeout(1 * time.Second)

    var ph mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
        if msg.Topic() == commandTopic && string(msg.Payload()) == "EXIT" {
            log.Printf("Got EXIT message on topic \"%s\"", aliveTopic)
            close(cdone)
        }
    }
    opts.SetDefaultPublishHandler(ph)

    c := mqtt.NewClient(opts)
    log.Println("Connecting MQTT client...")
    if token := c.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    log.Println("Connected MQTT client")

    publish5TestMessages(c)

    if token := c.Subscribe(commandTopic, 0, nil); token.Wait() && token.Error() != nil{
        panic(token.Error())
    }
    log.Printf("Send 'EXIT' to topic %s to exit\n", commandTopic)

    go func() {
        text := "ALIVE"
        log.Printf("Starting ALIVE Publishing \"%s\" on topic %s...\n", text, aliveTopic)
        for {
            token := c.Publish(aliveTopic, 0, false, text)
            token.Wait()
            time.Sleep(aliveFrequency)
        }
    }()
    select {
    case <-cdone:
        log.Println("Exiting MQTT ALIVE publishing routine...")
    }

    if token := c.Unsubscribe(aliveTopic); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    c.Disconnect(250)

    time.Sleep(1 * time.Second)
    log.Println("Finishing MQTT client gracefully...")
    wg.Done()
}
