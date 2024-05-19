package main

import (
	"fmt"
	"log"
)

type Employee struct {
	name string
	id string
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee %s with id %s", e.name, e.id)
}

type Manager struct {
	Employee
}


func main() {
	// Will directly call the String() method of Employee
	m := Manager{Employee{"Sam", "1234"}}
	log.Println(m)
}
