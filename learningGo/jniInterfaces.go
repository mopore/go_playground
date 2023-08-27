package main

import (
	"log"
	"fmt"
)

type Logic interface {
	Process(date string) string
}

type Poser interface {
	Pose()
}

type LogicProvider struct {}

func (lp LogicProvider) Process(data string) string {
	processed := fmt.Sprintf(">%s<", data)	
	return processed
}

func (lp LogicProvider) Pose() {
	log.Println("Posing")
}

type Client struct {
	L Logic
}

func (c Client) DoSomething(data string) string {
	return c.L.Process(data)
}

func main() {
	var i any
	i = "test"
	log.Println(i)

	c := Client{
		L: LogicProvider{},
	}
	log.Println(c.DoSomething("test"))

	// Check for second interface
	p, ok := c.L.(Poser)
	if ok {
		p.Pose()
	} else {
		log.Println("LogicProvider does not implement Poser")
	}
}

