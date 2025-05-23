package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	// sumup is a variadic function, so you can pass any number of int arguments.
	sum := sumup(1, 10, 15, 40, -5)    // Pass values directly
	anotherSum := sumup(1, numbers...) // Use ... to expand a slice into individual arguments

	fmt.Println(sum)        // Output: 61
	fmt.Println(anotherSum) // Output: 27
}

// sumup takes a starting value and any number of additional int arguments.
// The ...int syntax means numbers is a slice of int, and you can pass as many as you want.
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue // Start with the starting value

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}

/*
What you need to know about variadic functions in Go:
- Use ...type in the parameter list to accept any number of arguments of that type.
- Inside the function, the variadic parameter is a slice.
- You can pass a slice to a variadic function using ... (e.g., sumup(1, numbers...)).
- Only the last parameter in a function can be variadic.
- Useful for functions like fmt.Println, append, etc.

Example:
func printAll(values ...string) { ... }
printAll("a", "b", "c")
*/
