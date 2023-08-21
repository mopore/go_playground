package main

import (
	"io"
	"log"
	"os"
)


func getFile(filename string) (*os.File, func(), error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, nil, err
    }
    return f, func() {
        f.Close()
    }, nil
}

func main() {
    log.Printf("Simple Cat \n")
    log.Printf("----------- \n")

    if len(os.Args) != 2 {
        log.Printf("Usage: %s <filename>\n", os.Args[0])
        log.Fatal("Please specify a filename")
    }

    f, closer, err := getFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer closer()

    data := make ([]byte, 2048)
    for {
        count, err := f.Read(data)
        os.Stdout.Write(data[:count])

        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
    }
}
