package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func WriteFile(filename string, lines []string, append bool) {
	fmt.Println("Creating file: ", filename)
	var file *os.File
	var err error
	if append {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	} else {
		file, err = os.Create(filename)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		// _, err = fmt.Fprintln(file, line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadWholeFile(filename string) {
	fmt.Println("Reading file in a whole: ", filename)
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
}

func ReadStats(filename string) {
	fmt.Println("Reading file stats: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	isdir := stats.IsDir()
	fmt.Println("Is directory: ", isdir)
	fmt.Println("File size (bytes): ", stats.Size())
	modtime := stats.ModTime().Format("2006-01-02 15:04:05")
	fmt.Println("Modifed time: ", modtime)
}

func ReadByLine(filename string) {
	fmt.Println("Reading file line by line: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// splitted := line.Split(" ")
		fmt.Println(line)
	}
}

func ReadByWord(filename string) {
	fmt.Println("Reading file word by word: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint) {
	fmt.Println("Reading file byte by byte: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, size)
	for {
		total, err := file.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
			break
		}
		chunk := buf[:total]
		fmt.Print(string(chunk))
	}
}

func main() {
	const filename = "testfile.txt"

	lines := []string{
		"This is the first line",
		"This is the second line",
	}

	WriteFile(filename, lines, false)
	ReadStats(filename)
	ReadWholeFile(filename)
	// ReadByLine(filename)
	// ReadByWord(filename)
	// ReadByBytes(filename, 10)
}
