package main

import "fmt"

func main() {
	s := []int{1, 2, 5} // slice literal
	fmt.Println(s, len(s), cap(s))

	sl := make([]int, 3)
	fmt.Println(sl, len(sl), cap(sl))
	// [0 0 0] 3 3
	sl[0] = 5
	// [5 0 0] 3 3

	s2 := make([]int, 3, 5)
	fmt.Println(s2, len(s2), cap(s2))
	// [0 0 0] 3 5
	s2[0] = 8
	// [8 0 0] 3 5
	s2[3] = 7 //***ERROR
}

/*
The length is what defines how many elements exist in the slice.

The capacity is how many elements could potentially exist if you use append() to add more elements.
s2 = append(s2, 10)  // len now 4
*/
