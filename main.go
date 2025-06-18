package main

import "fmt"

func main() {
	var a int = 5

	fmt.Println(a)
}

/*
Separate Stack for Separate thread
---------------------------------------------
- software1 run
- process1 run (code, data, stack, ... in RAM)
- thread1, thread2, ....
- OS handles how thread will be executed

- all functions run on stack (stack frame)
- thread executes line by line stack code
- CU, ALU, Register set execute thread

- thread1 -> stack
- thread2 -> stack
- why separate stack?
	- different threads doing different tasks
	- in a single stack there are some functions for a thread
	- if I use same stack than we will loose those functions
	- we cannot use same stack for all functions from different thread
	- so different stack so that functions of different stack can stay

- for each thread, 8MB thread stack is allocated. (Linux)
- stack can be anywhere in RAM

- Main part of OS / Code is Kernel
- Kernel do most of the task like scheduling etc...
- Kernel decide which VCPU will be used for which process

- thread create
- decide which functions to run
- stack will be allocated for this thread
- thread close

- process does not track how many threads it has, where are the other stacks etc.
- process only tracks code, data, stack, heap and the main thread
- so other supporting threads for a single process are becoming independent
- Kernel tracks how many threads for a single process

- independent thread ask kernel to fetch info from code/data segment..
*/
