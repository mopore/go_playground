package slowy

import (
    "time"
    "fmt"
)

type SlowComplicatedParser interface {
    Parse() (string, error)
}

type Slowy struct {
    name string
}

func New(name string) *Slowy {
    time.Sleep(time.Second * 3)
    return &Slowy{name: name}
}

func (s Slowy) Parse() (string, error) {
    r := fmt.Sprintf("My name is %s", s.name)
    return r, nil
}

