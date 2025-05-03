package main

import "fmt"

func main() {
	s := []int{1, 2, 5} // slice literal
	fmt.Println(s, len(s), cap(s))
}

/*
2 Phase
-------
1. Compilation Phase - executible binary file create
========================

code segment
---
main().....

2. Execution Phase
========================
- load the code segment to RAM
code segment
---
main().....

data segment
---

stact
----
stack frame for main
   -- arrray create
   -- s (Pointer(1), Length(3), Capacity(3))
pop main
*/
