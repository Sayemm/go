package main

import "fmt"

func main() {
	// pointer or address of memory (ram)

	x := 20
	addr := &x // address of x

	fmt.Println(x)
	fmt.Println(addr)
	fmt.Println(*addr) // value at address addr

	*addr = 30

	fmt.Println(x)
}
