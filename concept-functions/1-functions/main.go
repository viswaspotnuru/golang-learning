package main

import "fmt"

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Pass the address of the slice to the function (pointer to slice)
	// Here, we are just passing the address, not modifying the original slice
	doubled := doubledNumbers(&numbers)

	fmt.Println("Original numbers:", numbers) // Original numbers: [1 2 3 4 5]
	fmt.Println("Doubled numbers:", doubled)  // Doubled numbers: [2 4 6 8 10]
}

// doubledNumbers takes a pointer to a slice of ints as input
// In this example, the pointer is not necessary because we are not modifying the original slice
// We are only reading from it and creating a new slice with doubled values
func doubledNumbers(numbers *[]int) []int {
	dNumbers := []int{} // New slice to store doubled values

	// *numbers dereferences the pointer to access the original slice's values
	for _, number := range *numbers {
		dNumbers = append(dNumbers, number*2)
	}

	// dNumbers is a new, independent slice (not sharing memory with the original)
	return dNumbers
}

/*
Pointers in Go:
- A pointer holds the memory address of a variable, not the value itself.
- Use & to get the address of a variable (e.g., &numbers).
- Use * to dereference a pointer and access the value at that address (e.g., *numbers).

When to use pointers:
- Use a pointer if you want a function to modify the original value.
- Use a value (no pointer) if you only need to read or don't need to change the original.

In this code:
- We pass a pointer, but only read from the original slice.
- The function creates and returns a new slice, so the pointer is not needed.
- numbers and doubled are two separate slices with different addresses.
*/
