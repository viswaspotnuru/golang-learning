package main

import (
	"fmt"
	"math"
)

// 1. THE STRUCT (The Blueprint)
// Represents an instance in memory.
type Vertex struct {
	X, Y float64
}

// 2. VALUE RECEIVER (Instance Receiver)
// - Memory: Go creates a full COPY of the 'v' instance.
// - Behavior: It cannot modify the original X or Y values.
// - Use Case: Use for "Read-Only" operations or small structs.
func (v Vertex) Abs() float64 {
	// v here is a local copy; changes to it die when the function ends.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 3. POINTER RECEIVER (* Receiver)
//   - Memory: Go passes only the MEMORY ADDRESS (8 bytes).
//   - Behavior: It modifies the original instance directly.
//   - Use Case: Use when you need to update state or for large structs
//     to avoid expensive copying.
func (v *Vertex) Scale(f float64) {
	// v here is a pointer; we are reaching into the original memory address.
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// --- INSTANCE CREATION ---
	// 'v' is an instance of Vertex allocated on the stack.
	v := Vertex{X: 3, Y: 4}

	fmt.Println("--- Initial State ---")
	fmt.Printf("Vertex: %+v | Abs: %f\n", v, v.Abs())

	// --- CALLING A POINTER METHOD ---
	// Even though 'v' is a value, Go automatically treats this
	// as (&v).Scale(10) behind the scenes.
	v.Scale(10)

	fmt.Println("\n--- After Scale(10) [Pointer Receiver] ---")
	fmt.Printf("Vertex: %+v (Original was modified!)\n", v)

	// --- MEMORY ADDRESS COMPARISON ---
	// You can see the address of the instance
	fmt.Printf("\nMemory Address of v: %p\n", &v)
}
