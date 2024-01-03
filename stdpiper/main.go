package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type NoInputError struct {
    location string
}

func NewInputError(location string) error {
    return NoInputError{location}
}

func (e NoInputError) Error() string {
    return fmt.Sprintf("%s: No input provided", e.location)
}

func parseInput() ([]string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, NewInputError("parseinput")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		noInputError := NewInputError("parseinput")
		return nil, fmt.Errorf("parseInput: error reading standard input: %w", noInputError)
	}

	return lines, nil
}

func main() {
	lines, err := parseInput()
	if err != nil {
        var customErr NoInputError
        if errors.As(err, &customErr) {
            fmt.Println("Please pipe input to this program")
            log.Fatalf("Could not parse input: %v", err)
        } else {
            log.Fatal(err)
        }
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
