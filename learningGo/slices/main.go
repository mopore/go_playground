package main

import (
	"fmt"
    "slices"
)

func describeSlice(name string, s []int) {
    fmt.Printf("name: %s\n", name)
    fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
    for _, v := range s {
        fmt.Printf("v: %d\n", v)
    }
}

func main() {
    s := []int{1, 2, 3, 4, 5}
    s = append(s, 6)
    describeSlice("Orignal", s)
    clone := make([]int, len(s))
    copy(clone, s)
    // Remove the second element
    clone = append(clone[:1], clone[2:]...)
    describeSlice("Clone without second", clone) 
    fmt.Println("--------------")
    clone = slices.Delete(clone, 1, 2)
    describeSlice("Clone without second", clone)
}
