package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const (
	filename = "input_example.json"
)

type Example struct {
	Name     string    `json:"name"`
	Weight   int       `json:"weight"`
	Birthday time.Time `json:"birthday"`
}

func toString(in Example) string {
	out, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func fromString(in string) Example {
	var example Example
	err := json.Unmarshal([]byte(in), &example)
	if err != nil {
		log.Fatal(err)
	}
	return example
}

func toFile(in Example, fn string) error {
	outFile, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = json.NewEncoder(outFile).Encode(in)
	if err != nil {
		return err
	}
	return nil
}

func fromFile(fn string) Example {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(f)
	}
	defer f.Close()

	var example Example
	err = json.NewDecoder(f).Decode(&example)
	if err != nil {
		log.Fatal(err)
	}
	return example
}

func main() {
	birthday, err := time.Parse("2006-01-02", "1980-01-08")
	if err != nil {
		log.Fatal(err)
	}
	exOriginal := Example{Name: "Julia", Weight: 59, Birthday: birthday}
	log.Printf("Original example as type: %v\n", exOriginal)

	exAsString := toString(exOriginal)
	log.Printf("JSON string from original: %v\n", exAsString)

	exFromString := fromString(exAsString)
	log.Printf("Example decoded/unmarshaled from string: %v\n", exFromString)

	err = toFile(exFromString, filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Example (from string) written to file")

	exFromFile := fromFile(filename)
	log.Printf("Example from file as type: %v\n", exFromFile)

	os.Remove(filename)
}
