package main

import "fmt"

func main() {
	// a := 10 // variable expression

	// Function Expression or Assign funtion in variable

	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(2, 3)
}
