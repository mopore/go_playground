package main

import (
	"errors"
	"fmt"
	"log"
)

var errorOriginal = errors.New("original error")

func errorThrower() error {
	return errorOriginal
}

func internalFunction() error {
	err := errorThrower()
	if err != nil {
		return fmt.Errorf("internalFunction: %w", err)
	}
	return nil
}

func main() {
    log.Println("Starting error test program")
    err := internalFunction()
	if err != nil {
		switch {
		case errors.Is(err, errorOriginal):
			log.Fatalf("main: need to act on orignal error: %v", err)
		default:
			log.Fatalf("Acting on unkown error: %v", err)
		}
	}
    log.Println("Error testing program finished unexpectedly without error")
}

