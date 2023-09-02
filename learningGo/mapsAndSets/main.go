package main

import "log"

func logOne(m map[string]int){
    i, ok := m["one"]
    if ok {
        log.Println("m[\"one\"] =", i)
    } else {
        log.Println("m[\"one\"] not found")
    }
}

func main(){
    log.Println("Testing maps")
    m := map[string]int{
        "one": 1,
        "two": 2,
    }
    logOne(m)
    delete(m, "one")
    logOne(m)

    // Simulating a set
    intSet := map[int]bool{}
    vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    for _, v := range vals {
        intSet[v] = true
    }
    log.Println("Has 5:", intSet[5])
    log.Println("Has 10:", intSet[10])
}
