package main

import (
	"fmt"
	"sync"
        "runlazyonce/slowy"
)

var parser slowy.SlowComplicatedParser
var once sync.Once

// Parse applies the SlowComplicatedParser.Parse() method but when called
// initialieses the parser only one.
func Parse() (string, error) {
    once.Do(func() {
        parser = initParser()
    })
    return parser.Parse()
}

// initParser is a helper function that initialises the parser.
// 
// Calling slowy.New() is a very slow operation.
func initParser() slowy.SlowComplicatedParser {
    fmt.Println("Initializing parser ...")
    s := slowy.New("John")
    return s
}

func main() {
    fmt.Println("Starting ...")
    fmt.Println("Calling Parse() ...")
    result1, err1 := Parse()
    if err1 != nil {
        panic(err1)
    }
    fmt.Println(result1)

    fmt.Println("Calling Parse() again ...")
    result2, err2 := Parse()
    if err2 != nil {
        panic(err2)
    }
    fmt.Println(result2)

    fmt.Println("Done.")
}
