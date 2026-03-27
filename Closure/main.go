package main

import "fmt"

func incrementor() func() int {
	// This variable is "captured" by the closure
	count := 0

	return func() int {
		count++ // The inner function modifies the outer variable
		return count
	}
}

func main() {
	// 'next' is now a closure that carries its own 'count' variable
	next := incrementor()

	fmt.Println(next()) // Output: 1
	fmt.Println(next()) // Output: 2

	// Creating a NEW incrementor starts a fresh count
	newNext := incrementor()
	fmt.Println(newNext()) // Output: 1
}
