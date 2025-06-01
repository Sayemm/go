package main

import (
	"fmt"
)

func add(x, y int) int {
	var result int
	result = x + y
	return result
}

func main() {
	var a int = 10
	var sum = add(a, 4)

	fmt.Println(sum)
}

/*
Computer 8 bits - means register set 8 bits - means - 000000000
Ram and HD will be 8 bits/1 byte as well
RAM - 0(Each will hold 8 bits) 1 2 3 4 5 6 7 8 .....

Computer 16 bits / 2 bytes
RAM - each cell will allocate 16 bits / 2 bytes (8 + 8)
    - 0 2 4 6 8 10

Computer 32 bits / 4 bytes
RAM - each cell allocate 32 bits / 4 bytes (8 + 8 + 8 + 8)
    - 0 4 8 12 16 20...

Computer 64 bits / 8 bytes
RAM - each cell allocate 64 bits / 8 bytes (8 x 8)
    - 0 8 16...

- Computer ON (32 bits -> 0 4 8 12 16 20...)
- OS will load to RAM from HD
- Control Unit will tell PC to point the first address (0)
- CU will fetch the instruction from the address that PC is indicating and put it to IR
- CU read IR and decode the 32 bit instruction (+ 3 5)
- EX: 3 will be on Accumulator Register (EAX), 5 on base register (EBX) and + on counter register (ECX)
- EX: CU instruct ALU to add and CU will put it to data register (DAX)
- PC value will increase and next instruction will be executed...
- Now computer is ON and OS is running..

- main.go will be saved to HD
- compile the code (go build main.go)
- binary file will be saved to HD
- Execute (./main)
	- OS will take the control
	- OS will create a PROCESS (new memory will be allocated)
		- CODE SEGMENT
			- func main()..
			- func add()...
		- DATA SEGMENT
		- STACK SEGMENT/FUNCTION FRAME - WILL CREATE FROM THE END
			- stack frame for main..
				- return address    ---80
				- old BP value (8) and BP value will change ---76
				- (a = 10) ----72
				-** SP will be 72 and BP will be 76
				==
				-** Now SP = 72 and BP = 76 AGAIN
				- sum = 14 ---68 -> S-> SP will be 68
				- print will find sum using BP - 8 = 68
				- pop sum -> SP will be 72
				- pop a   -> SP will be 76
				- SP = BP, stack is complete almost
				- BP will be old value = 8
				- return address means will go back to OS
			== stack frame for add...
				- y value = 4 ---68    -> SP will be 68
				- x value = 1- --64    -> SP will be 64
				- return address -- 60 -> SP will be 60
				- old BP value (76) and BP value will change ---56 -> SP will be 56
				- result = 0 -- 52 > SP will be 52
					- Computer will take help from BP to find X (BP + 8) and Y (BP + 12) and Result (BP - 4)
					- result = summation of X, Y
				- return result
					- pop 52 and SP will be 56
					- add stack pointer will return result to main stack pointer
					- BP will be now 76 and pop stack pointer
					- SP will be 72
			-** Now SP = 72 and BP = 76
		- HEAP
	- OS will give the CPU control to process
	- PC will point the first cell of the process


=> SP: When stack will change the SP value will change (increase/decrease)
=> BP:
	This is fixed in a certain stack frame,
	using BP we can access variables in a stack frame (+/-)
	when stack got popped BP will be the old BP value.

*/
