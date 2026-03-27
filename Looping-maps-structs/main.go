package main

import "fmt"

// 1. STRUCT: Defining the "Shape" of our data
type Furniture struct {
	Name  string
	Price float64
}

func main() {
	// --- DECLARATIONS ---

	// 2. ARRAY: Fixed size (3 slots). Memory is allocated immediately.
	var storageArray = [3]Furniture{
		{Name: "Couch", Price: 500},
		{Name: "Table", Price: 200},
	}

	// 3. SLICE: A dynamic window pointing to the array above.
	// We use 'make' to create a NEW slice with a pre-allocated capacity of 10.
	items := make([]Furniture, 0, 10)

	// 4. MAP: A fast lookup table (Key: Name, Value: Furniture Struct)
	catalog := make(map[string]Furniture)

	// --- LOGIC & OPERATIONS ---

	// Use 'append' with the '...' (unpack) operator to move array data into our slice
	// This takes the elements FROM the array and puts them INTO the slice.
	items = append(items, storageArray[:]...)

	// Add 5 more items to the slice (now we have 7 total)
	newItems := []Furniture{
		{Name: "Lamp", Price: 40},
		{Name: "Rug", Price: 150},
		{Name: "Chair", Price: 75},
		{Name: "Desk", Price: 300},
		{Name: "Mirror", Price: 60},
	}
	items = append(items, newItems...)

	// Populate the MAP for fast searching
	for _, f := range items {
		catalog[f.Name] = f
	}

	// --- OUTPUT ---

	fmt.Println("--- 1. The Slice (List of all items) ---")
	fmt.Printf("Total Items: %d | Capacity: %d\n", len(items), cap(items))

	fmt.Println("\n--- 2. The Map (Quick Lookup) ---")
	// Finding an item by name is O(1) time - instant!
	val, exists := catalog["Desk"]
	if exists {
		fmt.Printf("Found in Catalog: %s costs $%.2f\n", val.Name, val.Price)
	}
}
