package main

import (
	"fmt"
	"iter"
	"log"
)

func reverse[V any](s []V) iter.Seq[V] {
    return func(yield func(V) bool) {
        for i := len(s) - 1; i >= 0; i-- {
            if !yield(s[i]) {
                return
            }
        }
    }
}

func printAll[V any](s iter.Seq[V]) {
    fmt.Println("Printing all values")
    for v:= range s {
        fmt.Printf("%v\n", v)
    }
}


func main(){

    myMap := make(map[string]string)
    myMap["name"] = "John"
    myMap["age"] = "25"

    for key, value := range myMap {
        log.Println(key, value)
    }

    n := []int{1, 2, 3, 4, 5}
    printAll(reverse(n))

}
