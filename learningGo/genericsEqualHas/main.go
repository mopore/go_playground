package main

type Equalizer[T any] interface {
    Equal(T) bool
}

type Person struct {
    Name string
}

func (p Person) Equal(other Person) bool {
    return p.Name == other.Name
}

type Animal struct {
    Age int
}

func (a Animal) Equal(other Animal) bool {
    return a.Age == other.Age
}

// HasEqual returns true if s contains v.
// Passing 'T Equalizer[T]' as a type parameter ensures that each element of s
// can be compared to v.
func HasEqual[T Equalizer[T]](s []T, v T) bool {
    for _, e := range s {
        if e.Equal(v) {
            return true
        }
    }
    return false
}

func main() {
    people := []Person{
        {"Alice"},
        {"Bob"},
    }
    otherPeople := []Person{
        {"Charlie"},
        {"Dave"},
    }
    alice := Person{"Alice"}
    println(HasEqual(people, alice)) // prints true
    println(HasEqual(otherPeople, alice)) // prints false

    // The line below will not compile.
    // Interfaces are ignored when checking type parameters.
    // println(HasEqual([]Animal{Animal{1}}, alice))
}
