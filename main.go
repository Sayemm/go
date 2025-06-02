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
Context Switching | PCB | Concurrency
-------------------------------------

- Computer ON
- .....
- CU will start to execute instructions

- Software1, Software2, Software3  in RAM
- Double click (Start) - loads to RAM 1, 2, 3 - 3 process
- PC will point only one cell???!!! THEN HOW ARE 3 PROCESS RUNNING AT THE SAME TIME???

- MODERN COMPUTER CAN EXECUTE 10^8 INSTRUCTIONS IN A SEC (approx..)
- Old OS ran 3 process one by one, no profit in that..
- OS has a block name PCB (Process control block)

- Software1 (28...)-100 instructions, Software2 (52...) 200 instructions, Software3 (72...)300 instructions
- PC will be on 28
- OS will instruct to execute let's say 10 instructions on Software1
- Then OS will take PC to Software2 and PC will be 52
- OS will instruct to execute let's say 20 instructions on Software2
- Then OS will take PC to Software3 and PC will be 72
- OS will instruct to execute let's say 20 instructions on Software3
- .........
- THIS IS CALLED CONTEXT SWITCHING

- HOW PC knows where to go back of software1/2/3 to execute the rest instructions and what about past registers values?
- PCB comes to play here, IT TRACKS ALL THE PROCESS
- State (all registers valueS (SP, BP, IR, PC, AL, BL, CL, DL...), PID (process ID), ....)
- This state is saved to PCB and goes to 2nd process and check for process 2 in PCB.....
- OUR BRAIN THINKS EVERYTHING IS HAPPENNING AT THE SAME TIME (1/10) BUT NOT..
	- THIS IS CONCURRENCY - EVERYTHING IS NOT HAPPENNING AT THE SAME TIME, HAPPENNING BY SWITCHING CONTEXT, HANDLING LOTS OF PROCESS AT THE SAME TIME
	- IT GIVES A FEEL THAT, EVERYTHING IS HAPPENNING AT THE SAME TIME BUT NOT
*/
