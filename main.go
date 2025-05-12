package main

func main() {

}

/*

CPU
---
Processing Unit
	Control Unit -
	ALU - + - * / & | !
Register Set
	- PC - Program Counter (Pointer Register)
	     - Point on RAM which portion to execute
	- Instruction Register (Load from the RAM fetched by CU)
	- SP - Stack Pointer - end of stack frame
	- BP - Base Pointer - anytwhere in that stack frame
	-> General Purpose Registers..
	- AL - Accumulator Register (8 bit) (16 bit -> AX [AH (lower) + AL(higher)]) (32 - EAX) (64 - RAX)
	- BL - Base Register
	- CL - Counter Register
	- DL - Data Register
RAM
---

HD
--

* PC points to RAM first cell when OS is load to RAM from HD
* CU fetch the instruction from RAM indicated by PC and put it to IR
* Increase PC value by 1
* CU reads IR and decode the instruction
* CU orders ALU to do the operations

go
--
* file on HD
* compile -> go build -> binary executable file -> save HD
* execution -> binary file will load to RAM (code segment, data segment, stack, heap)

Process / Virtual or Logical Computer (cz Process has everything that a computer has)
--------------------------------------
-> execute the first to last line of instruction, complete task using..
- code segment - data segment - stack - heap
- CPU
*/
