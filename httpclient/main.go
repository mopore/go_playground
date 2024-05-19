package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type option func(*http.Client)

func WithTimeout(timeout time.Duration) option {
	return func(c *http.Client) {
		c.Timeout = timeout
	}
}

func New(opts ...option) *http.Client {
	c := &http.Client{}
	c.Timeout = 30 * time.Second
	for _, opt := range opts {
		opt(c)

	}
	return c
}

func GetGoogle() {
	c := New(
		WithTimeout(10 * time.Second),
	)
	res, err := c.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}

func parseJsonFromGet(){
    c := New(
        WithTimeout(10 * time.Second),
    )
    req, err := http.NewRequestWithContext(
        context.Background(), 
        http.MethodGet, 
        "https://jsonplaceholder.typicode.com/posts/1",
        nil,
    )
    req.Header.Set("Accept", "application/json")
    req.Header.Add("X-My-Client", "Learning Go")
    res, err := c.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != http.StatusOK {
        log.Fatal(res.Status)
    }
    fmt.Println(res.Header.Get("Content-Type"))

    var data struct {
        UserID int `json:"userId"`
        ID int `json:"id"`
        Title string `json:"title"`
        Completed bool `json:"completed"`
    }
    err = json.NewDecoder(res.Body).Decode(&data)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", data)
}

func main() {
    parseJsonFromGet()
}
