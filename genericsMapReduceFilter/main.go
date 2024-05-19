package main

import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T, U any](s []T, initializer U, f func(U, T) U) U {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}


func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
    words := []string{"hello", "world!"}
    filtered := Filter(words, func(s string) bool {
        return s == "hello"
    })
    fmt.Printf("%v\n", filtered)
    
    wordLenghts := Map(words, func(s string) int {
        return len(s)
    })
    fmt.Printf("%v\n", wordLenghts)

    sum := Reduce(wordLenghts, 0, func(a, i int) int {
        return a + i
    })
    fmt.Printf("%v\n", sum)
}

