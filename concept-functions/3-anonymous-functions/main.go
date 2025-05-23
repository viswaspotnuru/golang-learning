package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Here we use an anonymous function as the transform argument.
	// Anonymous functions are functions without a name, defined inline.
	// They are useful when you need a simple, one-off function and don't want to declare it separately.
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2 // This function doubles each number.
	})

	fmt.Println(transformed) // Output: [2 4 6]
}

// transformNumbers takes a pointer to a slice and a function as arguments.
// The function argument can be any function matching the signature (int) int,
// including anonymous functions.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

/*
Why use anonymous functions?
- They allow you to define custom logic inline, making code concise and readable.
- Useful for simple transformations or callbacks used only once.
- Avoids polluting your code with many small, named functions.
*/
