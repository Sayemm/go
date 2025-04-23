package main

import "fmt"

func main() {
	arr := [6]string{"THIS", "IS", "A", "GO", "INTERVIEW", "QUESTION"}

	fmt.Println(arr)
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
   -- arr
   -- print
pop main
*/
