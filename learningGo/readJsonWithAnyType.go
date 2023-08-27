package main

import (
	"encoding/json"
	"os"
	"log"
)


func main() {
	data := map[string]any{}
	filecontent, err := os.ReadFile("file.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(filecontent, &data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}
