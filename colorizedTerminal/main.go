package main

import (
    "fmt"
    "github.com/gookit/color"
)

func main() {
    color.Bold.Println("Bold line")
    color.Red.Println("Simple to use color")
    color.Green.Print("Simple to use color\n")
    color.Cyan.Printf("Simple to use %s\n", "color")
    color.Yellow.Printf("Simple to use %s\n", "color")

    // use like func
    red := color.FgRed.Render
    green := color.FgGreen.Render

    fmt.Printf("%s line %s library\n", red("Command"), green("color"))
}


