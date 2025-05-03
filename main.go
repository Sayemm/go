package main

import "fmt"

func main() {
	arr := [6]string{"THIS", "IS", "A", "GO", "INTERVIEW", "QUESTION"}

	fmt.Println(arr)

	s := arr[1:4] //points arr ["IS", 3, 5]
	fmt.Println(s)

	s1 := s[1:2] // untimately points arr as well ["A", 1, 4]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

// Slice Maintains 3 things: Pointer(1), Length(3), Capacity(5 -> starts at 1 to end)

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
   -- arr
   ** s (ptr, len, cap)
   ** s1 (ptr, len, cap)
   -- print
pop main
*/
