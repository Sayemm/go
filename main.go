package main

import "fmt"

func add(a int, b int) { // parameters (a, b) - when we receive something on a function
	fmt.Println(a + b)
}

func main() {
	add(2, 3) // arguments (2, 3) - when we pass value to a fuction
}
