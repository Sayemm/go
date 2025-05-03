package main

import "fmt"

func main() {
	var s []int // empty or nil slice / ptr = nil
	fmt.Println(s, len(s), cap(s))
	// [] 0 0

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
	// [1] 1 1

	s = append(s, 2)
	fmt.Println(s, len(s), cap(s))
	// [1, 2] 2 2
}

/*
s = append(s, 1)
--
ptr = nil so
array create on main stack frame:
	array will stay on heap as the array will pop from stack fram of main
*/
