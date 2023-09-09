package main

import (
	"fmt"
	"math"
)

// Currency is a type that represents a currency.
// A variation of int or int64 is used to represent the currency amount.
// The ISO4127Code method returns the ISO 4127 code for the currency.
// The Decimal method returns the number of decimal places for the currency.
type Currency interface {
    ~int | ~int64
    ISO4127Code() string
    Decimal() int
}

type NZD int64

func (n NZD) ISO4127Code() string {
    return "NZD"
}

func (n NZD) Decimal() int {
    return 2
}

func PrintBalance[C Currency](b C) {
    balance := float64(b) / math.Pow10(b.Decimal())
    fmt.Printf("%s %.2f\n", b.ISO4127Code(), balance)
}

func main() {
    PrintBalance(NZD(100))
}
