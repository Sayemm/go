package main

import "fmt"

func main() {
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3}
	d := [...]int{1, 4: 5, 6}

	fmt.Println(len(a), len(b), len(c))
	fmt.Println(d)

	for i := range 3 {
		for j := range 3 {
			fmt.Println(i + j)
		}
	}
}

/*
Threads
---------------------------------------------

- Computer ON
- Load OS portion from HD to RAM
- now this OS will control everything (computer hardware - cpu, ram, ...)

- 2 software on
- load on the RAM from HD
- 2 process create on RAM
- if 1 VCPU then 2 process will run one by one using context switching.

- process 1 had (code (say 10 lines), data segment, stack, heap)
- PC will point the first line of code
- CU will take this line and put it to IR
- CU will decode the code line and give it to ALU to
- next PC....

- this PC value is manipulated by OS

* thread is a unit that is executed.
* software 1 has a default thread
* when we are saying executing process, we are executing thread
* thread is a virtual process

* process can create multiple thread

* backend server
	- single process this server is running on, single thread
	- 100 request to this backend server
	- if the process executes 100 request using 1 thread it will take a lots of time
	* this process will now create 100 more thread and each thread will take 1 request
	* 100 request will now execute like a single request

* software 1 say music player, single process but doing multiple things
	- showing music time, play, pause, list, remaining time,.....
	- multiple thread doing these jobs
* code segment has multiple lines, says 10 lines
* thread 1 is responsible for some lines, thread 2 is responsible for some lines, ....
* thread can access the code segment and responsible for different parts.
* thread can execute the code at the same time.

-----
* OS will keep the first line of code in PC
* now 1 process ie. 1 thread
* PC -> IR -> ALU....
* code seg will feel that the code is being executed by that thread
* will create another thread for another task of that process
* now both thread are executing the code at the same time.

* how both thread are executing at the same time? PC can point to only one line?
	- PC will point one thread (that was pointing the portion of a code) at a time
	- CONTEXT SWITCHING WILL HAPPEN WITHING THREADS NOW
	- Threads past memory will be saved somewhere

* threads context switching take less time than process context swithcing
 - because for process context switching, need to save the full state using PCB

* 2 software on, 2 process (p1, p2)
- 1 VCPU
- PC will point p1 code, thread 1.., PCB
- PC will point p2 code, thread 1.., PCB
- PC will point p1 code, pc->thread 1, pc->thread 2.., PCB
- ..........


* MULTIPLE 2 core, 4 VCPU
- 4 PC, 4 ALU, .....
- one VCPU will process p1
- another VCPU will process p2
	PARALLALISM
- no context switching between processes
- but context switching will happen between threads in a single process

* THEN ANOTHER 2 VCPU DOING NOTHING??!!
- Lets' say total 100+ threads for all the processes
- OS will use all the VCPU to run 100+ threads
- Context switching, Parallelism keep happenning depending on the situation


* RAM is different for 2 processes
- but RAM is same for multiple threads
- so it's easy to send data/commmunicate between threads which is hard between processes (PCB)
- means threads can use same data but processes cannot
*/
