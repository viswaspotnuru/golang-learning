package main

import "fmt"

func main() {
	// Array: length and capacity are both 4 because arrays in Go have a fixed size.
	arr := [4]int{10, 20, 30, 40}
	fmt.Println("Array:", arr)
	fmt.Println("Array length:", len(arr))   // len(arr) returns 4 (number of elements)
	fmt.Println("Array capacity:", cap(arr)) // cap(arr) returns 4 (fixed size of the array)

	// Slice: created with length 3 and capacity 5.
	// Length is the number of initialized elements, capacity is the total space allocated.
	slice := make([]int, 3, 5)
	fmt.Println("Slice:", slice)
	fmt.Println("Slice length:", len(slice))   // 3 (number of initialized elements)
	fmt.Println("Slice capacity:", cap(slice)) // 5 (total allocated space)

	// Appending two elements increases the length to 5, which matches the capacity.
	slice = append(slice, 1, 2)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Slice length:", len(slice))   // 5 (now using all allocated space)
	fmt.Println("Slice capacity:", cap(slice)) // 5 (capacity remains the same)

	// Appending one more element exceeds the current capacity.
	// Go automatically increases the capacity (usually doubles it).
	slice = append(slice, 3)
	fmt.Println("Slice after more append:", slice)
	fmt.Println("Slice length:", len(slice))   // 6 (number of elements after append)
	fmt.Println("Slice capacity:", cap(slice)) // Capacity increases (implementation-dependent)

	// --- Dynamic List Example with Slices ---
	// Create an empty slice of integers
	var dynamicList []int

	// Append elements dynamically
	dynamicList = append(dynamicList, 10)
	dynamicList = append(dynamicList, 20, 30)
	fmt.Println("Dynamic List:", dynamicList) // [10 20 30]

	// You can append another slice as well
	more := []int{40, 50}
	dynamicList = append(dynamicList, more...)
	fmt.Println("After appending more:", dynamicList) // [10 20 30 40 50]

	// Length and capacity grow as needed
	fmt.Println("Dynamic List Length:", len(dynamicList))   // 5
	fmt.Println("Dynamic List Capacity:", cap(dynamicList)) // Capacity increases automatically
}
