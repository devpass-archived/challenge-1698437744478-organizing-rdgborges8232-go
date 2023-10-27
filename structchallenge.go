package main

import (
    "fmt"
)

type Person struct {
    Name    string
    Age     int
    Country string
}

func main() {
    // Add your solution here!
    person := Person{
        Name:    "John Doe",
        Age:     30,
        Country: "USA",
    }

    fmt.Println(person)
}
