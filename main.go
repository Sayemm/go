package main

import "fmt"

func a() {
	i := 0
	fmt.Println("First i = ", i)

	defer fmt.Println("Second = ", i)

	i++
	fmt.Println("Third = ", i)

	defer fmt.Println("Fourth i = ", i)
}

func main() {
	a()
}

/*
Defer
---------------------------------------------
- defer does not execute the line immediately
- evaluated / store somewhere immediately but execution will be later


2 Phase:
	1. Compilation Phase (compile time)
	2. Execution phase (runtime)

*** Compile Phase ***

--- Code Segment ---
a := func () {...}
main = func () {...}

go run main.go => compile it => main => ./main
go build main.go => compile it => main


*** Execution Phase ***
./main => EXECUTION PHASE WILL START..

- BINARY FILE from HD code segement will be copied to RAM code segment

--- Code Segment ---
a := func () {...}
main = func () {...}

--- Code Segment ---

--- Stack ---
- stack frame for main -> code base of main will be executed now
- a will be found in code segment (not in main stack frame, not in data segemet, gotcha in data segment)
- stack frame for a
	- i = 0
	=> print First i = 0
	* defer => fmt.Println("Second = ", i) will be stored somewhere (i = 0)
	- i = 1
	=> print Third i = 1
	- * defer => fmt.Println("Fourth = ", i) will be stored somewhere (i = 1) on top of other one (like a STACK)
	- return, so stack frame should be popped
	- but before popping the stack frame, go runtime will call the function => fmt.Println("Fourth = ", i)
	- stack frame for Println
	=> print Fourth i = 1
	- pop Println stack frame
	- go runtime will call the function => fmt.Println("Second = ", i)
	- stack frame for Println
	=> print Second i = 0
	- pop Println stack frame
- pop a stack frame
- pop main stack frame
- Process is does / RAM will be cleared
*/
