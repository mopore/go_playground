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
	L Logic  // has to come with a property "L" following the "Logic" interface
}

func (c Client) DoSomething(data string) string {
	return c.L.Process(data)  // using its property following the "Logic" interface
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
	if p, ok := c.L.(Poser); ok {
		p.Pose()
	} else {
		log.Println("LogicProvider does not implement Poser")
	}
}

