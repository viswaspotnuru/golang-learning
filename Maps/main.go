package main

import "fmt"

func main() {
	// --- Maps Example ---
	// A map is a collection of key-value pairs. Keys must be unique.
	// Syntax: map[KeyType]ValueType

	// Create a map with string keys and int values
	ages := make(map[string]int)

	// Add key-value pairs
	ages["Alice"] = 25
	ages["Bob"] = 30

	// Retrieve a value by key
	fmt.Println("Alice's age:", ages["Alice"])

	// Check if a key exists
	age, exists := ages["Charlie"]
	if exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie is not in the map.")
	}

	// Iterate over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// --- Structs Example ---
	// A struct is a collection of fields. Field names are fixed and known at compile time.

	type Person struct {
		Name string
		Age  int
	}

	// Create a struct instance
	alice := Person{Name: "Alice", Age: 25}
	bob := Person{Name: "Bob", Age: 30}

	// Access struct fields
	fmt.Println("Alice's age (struct):", alice.Age)
	fmt.Println("Bob's age (struct):", bob.Age)

	// --- Difference between Maps and Structs ---
	// Maps:
	//   - Use dynamic keys (can add/remove keys at runtime)
	//   - All values have the same type
	//   - Useful for unknown or changing sets of keys
	// Structs:
	//   - Use fixed field names (known at compile time)
	//   - Each field can have a different type
	//   - Useful for representing objects with a known structure
}
