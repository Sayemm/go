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

RAM
---

HD
--

* PC points to RAM first cell when OS is load to RAM from HD
* CU fetch the instruction from RAM indicated by PC and put it to IR
* Increase PC value by 1
* CU reads IR and decode the instruction
* CU orders ALU to do the operations

*/
