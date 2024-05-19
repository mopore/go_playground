package main

import (
	"log"
	"github.com/mopore/go-test-module/testmodule"
)

func main() {
    log.Println("Hello World")
    testValue := testmodule.TestFunction()
    log.Printf("Test Value: %s", testValue)
}
