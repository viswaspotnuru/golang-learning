package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// createTransformer returns a closure that multiplies by the given factor.
	double := createTransformer(2) // double is now a function that multiplies by 2
	triple := createTransformer(3) // triple is now a function that multiplies by 3

	// Using an anonymous function directly as a transformer (doubles each number)
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	// Using closures returned by createTransformer
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed) // [2 4 6]
	fmt.Println(doubled)     // [2 4 6]
	fmt.Println(tripled)     // [3 6 9]
}

// transformNumbers applies the transform function to each element in the slice.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// createTransformer returns a closure.
// The returned function "remembers" the value of factor even after createTransformer has finished executing.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor // factor is captured from the outer scope
	}
}

/*
What is a closure?
- A closure is a function value that references variables from outside its own body.
- The function may access and modify the variables even after the outer function has finished executing.

Why are closures useful?
- They allow you to "capture" and use context (like the factor) when creating functions dynamically.
- Useful for creating configurable or stateful functions.

Tricky points about closures in Go:
- Closures capture variables by reference, not by value. If the captured variable changes after the closure is created, the closure sees the updated value.
- If you create closures inside a loop, be careful: all closures may reference the same variable, leading to unexpected results.

Example of a tricky closure in a loop:
    funcs := []func(){}
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() { fmt.Println(i) })
    }
    // All funcs will print 3, not 0, 1, 2, because i is captured by reference.

To avoid this, capture the variable inside the loop:
    for i := 0; i < 3; i++ {
        v := i
        funcs = append(funcs, func() { fmt.Println(v) })
    }
*/
