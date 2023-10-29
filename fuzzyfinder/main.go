package main

import (
	"fmt"
	"log"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)


func main() {
    type YasmItem struct {
        Skillname string
        ID   int
    }

    var items = []YasmItem{
        { Skillname: "foo", ID: 1 },
        { Skillname: "bar", ID: 2 },
        { Skillname: "Foo3", ID: 3 },
    }

    // i, err := fuzzyfinder.Find(
    //     items,
    //     func(i int) string {
    //         return items[i].Name
    //     },
    // )

    i, err := fuzzyfinder.Find(
        items,
        func(i int) string {
            return items[i].Skillname
        },
        fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
            if i == -1 {
                return ""
            }
            return fmt.Sprintf("%s (id=%d)", items[i].Skillname, items[i].ID)
        }),
    )

    if err != nil { 
        log.Fatal(err)
    }
    item := items[i]
    log.Printf("selected: %s", item.Skillname)

}
