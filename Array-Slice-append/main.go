package main

import "fmt"

func main() {
	// Array of size 3
	storage := [3]int{10, 20, 30}

	// Slice of the first 2 elements.
	// Capacity is 3 (it can see to the end of the array).
	mySlice := storage[0:2]

	fmt.Printf("BEFORE: Array Addr: %p | Slice Addr: %p\n", &storage[0], &mySlice[0])
	// They will be IDENTICAL.

	// This append stays WITHIN capacity (3rd slot of array)
	mySlice = append(mySlice, 40)
	fmt.Printf("DURING: Array Addr: %p | Slice Addr: %p\n", &storage[0], &mySlice[0])
	// Still IDENTICAL. storage[2] is now 40.

	// THIS APPEND EXCEEDS THE ARRAY CAPACITY (4th element)
	mySlice = append(mySlice, 50)

	fmt.Printf("AFTER:  Array Addr: %p | Slice Addr: %p\n", &storage[0], &mySlice[0])
	// They are now DIFFERENT! The slice has "teleported" to a new home.
}
