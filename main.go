package main

import "fmt"

// Standard or Named Function
func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	add(3, 4)

	// Anonymous function
	// Immediately Invoked Function Expression (IIFE)
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(5, 7)
}

// Init Function - Executes before main function
func init() {
	fmt.Println("First function - I will be executed first")
}
