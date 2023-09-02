package main

import (
    "fmt"
    "log"
)

func errorThrower() error {
	return fmt.Errorf("original error")
}

func internalFunction() error {
	err := errorThrower()
	if err != nil {
		return fmt.Errorf("internal function: %w", err)
	}
	return nil
}

func main() {
    log.Println("Starting error test program")
    err := internalFunction()
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Error testing program finished unexpectedly without error")
}

