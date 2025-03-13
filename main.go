package main

import "fmt"

var (
	a2 = [3]string{"I", "Love", "You"}
)

func main() {
	var arr [2]int // by default 0
	anr := [2]int{3, 6}

	arr[1] = 6
	arr[0] = 5

	fmt.Println(arr)
	fmt.Println(anr)

	for index, num := range arr {
		fmt.Println(index, num)
	}
}

/*
Code Segment
---
main = func() {...}

data seg
-------
a2

Stack
-----
stack frame for main
2 place for arr
*/
