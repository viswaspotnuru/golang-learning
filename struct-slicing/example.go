package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. THE BACKING ARRAY (Fixed size, allocated on stack/heap)
	// This is a fixed block of memory for exactly 5 structs.
	var backingArray = [5]Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "David", Age: 28},
		{Name: "Eve", Age: 22},
	}

	// 2. THE SLICE (A "Window" into the Array)
	// Points to index 1 (Bob) and has a length of 2.
	sliceWindow := backingArray[1:3]

	// 3. THE POINTER SLICE (Slice of memory addresses)
	// Instead of storing the data, we store the "address" of the data.
	pointerSlice := []*Person{
		&backingArray[0], // Address of Alice
		&backingArray[4], // Address of Eve
	}

	fmt.Println("--- Before Modifications ---")
	fmt.Printf("Slice Window: %v\n", sliceWindow)
	fmt.Printf("Pointer Slice (Alice's Name): %s\n", pointerSlice[0].Name)

	// 4. MODIFICATION DEMO
	// Changing the pointer slice affects the original backing array.
	pointerSlice[0].Name = "Alice (Updated via Pointer)"

	// Changing the slice window also affects the original backing array.
	sliceWindow[0].Name = "Bob (Updated via Slice)"

	fmt.Println("\n--- After Modifications ---")
	fmt.Printf("Original Backing Array[0]: %s\n", backingArray[0].Name)
	fmt.Printf("Original Backing Array[1]: %s\n", backingArray[1].Name)
}
