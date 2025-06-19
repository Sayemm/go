package main

import "fmt"

func sum(a int, b int) (result int) {
	result = a + b
	return
}
func main() {
	res := sum(3, 4)
	fmt.Println(res)
}

/*
Named Return Value
-------------------
*/
