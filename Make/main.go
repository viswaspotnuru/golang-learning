package main

import "fmt"

func main() {
	// --- Using make() ---
	// make() is used to initialize slices, maps, and channels with optional size/capacity.

	// Using make to create a slice with length 3 and capacity 5
	slice := make([]int, 3, 5)
	fmt.Println("Slice with make:", slice)     // [0 0 0]
	fmt.Println("Slice length:", len(slice))   // 3
	fmt.Println("Slice capacity:", cap(slice)) // 5

	// Using make to create a map
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println("Map with make:", m) // map[one:1]

	// --- Normal Initialization ---
	// Slices: using a slice literal
	sliceLiteral := []int{1, 2, 3}
	fmt.Println("Slice literal:", sliceLiteral) // [1 2 3]

	// Maps: using a map literal
	mapLiteral := map[string]int{"two": 2, "three": 3}
	fmt.Println("Map literal:", mapLiteral) // map[two:2 three:3]

	// --- Differences ---
	// - make() allocates and initializes slices, maps, or channels, but does not set element values (except zero values).
	// - Literals allow you to specify initial values directly.
	// - make() is required when you want to specify capacity (for slices) or create an empty map/channel before adding elements.
}
