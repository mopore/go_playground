package main

import (
	"log"
	"net/http"
	"time"
)

type ClientOption func(*http.Client)

func New(opts ...ClientOption) *http.Client {
	c := &http.Client{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithTimeout(timeout int) ClientOption {
	return func(c *http.Client) {
		c.Timeout = time.Duration(timeout) * time.Second
	}
}

func main() {
    c := New(
        WithTimeout(10),
    )
    resp, err := c.Get("https://www.google.com")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%v\n", resp.Status)
}
