package main

import "fmt"

var (
	a = 10
)

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	add(5, 4)
	add(a, 6)
}

func init() {
	fmt.Println("Hello")
}

/*
code segment - data segment - stack - heap..(GC)

-> GC-Gurbage collector manages the heap
-> Global memory goes to data segment
-> All functions goes to code segment
-> When a function is called it takes a space in stack called stack frame where
   all the function var and etc take memory.

data segment - a
code segment - add, main, init
Stack - init (stack frame) -> pop
        main (stack frame) -> add (stack frame) -> pop add -> add (stack frame) -> pop add -> pop main

ALL CLEAR (code segment - data segment - stack - heap)
*/
