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
Concurrency (handling) Vs Parallelism (doing)
---------------------------------------------
- core i 3 CPU (guess)
	- Let's say 3 core in the CPU
	- core: 2 virtual/logical CPU (ALU, CU, Register Set)
	- 3 x 2 Virtual CPU in Core i 3

- Computer ON
- OS load to RAM
- Software1, Software2 run
- If Virtual CPU is available them separte VCPU for seperate task/software, don't need context switching
	- This is parallelism, context switching is not needed here.

- 3 process, 6 VCPU - Parallelism
- 6 process, 6 VCPU - Parallelism
- 7 process, 6 VCPU - 1->5 Parallelism, 6 & 7 will be Concurrency
                    - say process 1 is done then 6 & 7 will be done by separate VCPU ie Parallelism will be back.

- Cons of Context Switching
	- p1, p2, p3, only one VCPU -> Concurrency
	- PCB will get the state of process -> takes time to do these but this time is not needed in parallelism

-> INTEL CALLS VCPU/LCPU THREAD, BUT OS THREAD IS DIFFERENT

-> LETS SAY 1st process (10 minutes), 2nd process (2 miinutes)
	- if context switching takes 3 minutes then it will take 15 minutes to complete all the tasks (10 + 2 + 3)
	- without context switching it would take 12 minutes
	- so it depends
	- OS is smart and it uses different scheduling algorithm to do these tasks
*/
