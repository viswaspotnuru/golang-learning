package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a slice of Person structs
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// Print the slice
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}