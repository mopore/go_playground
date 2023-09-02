package main

import (
	"errors"
	"log"
)

type JniError struct {
	code int
}

func (e JniError) Error() string {
	return "jni error example"
}

func (e JniError) New(code int) JniError {
	return JniError{code}
}

func throwSimpleError() error {
	return errors.New("error example")
}

func throwCustomError() error {
	return JniError{}.New(99)
}


func main() {
	simpError := throwSimpleError()
	if simpError != nil {
		log.Printf("Error from throwSimpleError: %v\n", simpError)
	}
	customErr := throwCustomError()
	if customErr != nil {
		log.Printf("Error from throwCustomError: %v\n", customErr)
		log.Println("Checking if error is JniError")
		
		var jniErr JniError
		if errors.As(customErr, &jniErr) {
			log.Println("JniError code: ", jniErr.code)
		} else {
			log.Println("Error is not JniError")
		}
	}
	log.Println("Error testing program finished unexpectedly without error")
}
