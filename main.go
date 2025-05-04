package main

import "fmt"

// VARIADIC FUNCTION
func print(numbers ...int) { // numbers is a slice
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(4, 5, 6, 7, 8)
	print(4, 5, 6, 7, 8, 9, 10)
}
