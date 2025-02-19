package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, 4)
}

func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("INIT FUNC")
}

/*
1. Compilation Phase -> binary file create (01)
2. Executino Phase -> line by line execute

Code segment - Fixed ie. not possible to change - a, acll, add:call, main, init
Data segment - global var - p
Stack - when func runs it takes a place to stack frame - local scope
*/
