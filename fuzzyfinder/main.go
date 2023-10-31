package main

import (
	"log"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func selectWithPrevWindow() int {

    type YasmItem struct {
        Skillname string
        ID   int
    }

    var items = []YasmItem{
        { Skillname: "foo", ID: 1 },
        { Skillname: "bar", ID: 2 },
        { Skillname: "Foo3", ID: 3 },
    }

    i, err := fuzzyfinder.Find(
        items,
        func(i int) string {
            return items[i].Skillname
        },
        fuzzyfinder.WithCursorPosition(fuzzyfinder.CursorPositionTop),
        fuzzyfinder.WithPromptString("Select Skill: "),
    )

    if err != nil { 
        log.Println("No selection. Setting to 0")
        return 0
    }
    item := items[i]
    return item.ID

}

func selectSimple() string {

    options := []string{
        "eins",
        "zwei",
        "drei",
    }

    i, err := fuzzyfinder.Find(
        options,
        func(i int) string {
            return options[i]
        },
        fuzzyfinder.WithCursorPosition(fuzzyfinder.CursorPositionTop),
        fuzzyfinder.WithPromptString("Select Option: "),
    )

    if err != nil { 
        log.Fatal(err)
    }
    return options[i]

}

func main() {
    skillid := selectWithPrevWindow()
    log.Printf("selected skill ID: %d", skillid)
    option := selectSimple()
    log.Printf("selected option: %s", option)
}
