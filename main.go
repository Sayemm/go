package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int) {
	fmt.Println("Hello", num)
}

func main() {
	fmt.Println("Hello")
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(4)

	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second) // main go routrine will be done -> go runtime will be done ->process will terminate
	// 5s e e jara execute hoite parbe hobe naile nai
}

/*
Go Routine
----------
- go routine -> virtual thread (why virtual? works like thread but not thread)
- thread -> virtual process
- process -> virtual computer


Compilation Phase - go build main.go - main file
==============================
- binary executible file on HD

Code Segment
------------
p = 11
printHello()..
main()...


Execution Phase - ./main - binary file will load to the RAM
========================
* Process will be created and 4 things will be in that process
* BUT INITIALLY THERE WILL BE ONLY CODE SEGMENT - Binary will be loaded in Code segment

* Initially there is a thread
* Kernel of OS will start to execute the thread
		- CPU executes OS (OS has kernel)
		- Kernel tells CPU to execute thread
* thread executes stack
* and feels like process is executing...

* Process - Virtual Computer - when computer starts OS starts at first
* Go Runtime - mini OS will run first

- CODE SEGMENT
p = 11
printHello()..
main()...

- DATA SEGMENT
a = 10

- STACK
Stack frame for main. (WRONGGGGGGGGGGG)


- HEAP (dynamic - will change)


* GO RUNTIME itself is a virtual OS
* when go process starts go runtime starts and takes place in stack
* every process has a main thread and main stack (go runtime stack)
* main thread executes main stack or go runtime

* What does go runtime do?
	- go runtime initializes some things
		- go routine scheduler ini
		- heap allocator ini
		- GC initialize
		- logical processor create
			- BASED ON CPU
			- cpu has 4 virtual processor, go runtime will create 4 logical processor

* OS will create 4 separate OS thread - anywhere in memory (4 - stack for each thread)
* now total 5 thread in that process including main thread
* OS kernel will track 5 threads
* Go runtime creates go routine scheduler
	- 10 go routine, 4 mini thread
	- 			  3 3 2 2


* go runtime - mini OS - WILL CREATE 1 main GO ROUTINE/mini thread/virtual thread/logial thread AT FIRST
	- 2KB space for 1 go routine in HEAP!!
	- STACK FRAME FOR MAIN in 2KB
		- hello print
	- this is running by one Virtual Processor

	* another go routine in heap - 2KB stack outside main go routine
		- printHello stack frame
    - this go routine running by another virtual processor


*

ফিজিক্যাল cpu এর ভিতর থাকে core,
core এর ভিতর থাকে লজিক্যাল processor,
logical processor রান করে Operating system কে,
os এর kernel রান করে Go এর process কে,
Go র process এর ভিতর থাকে main thread, সে রান করে go run time কে,
go run time initialize করে virtual processor কে,
যেটা ম্যাপ হয় রিয়্যাল os thread এর সাথে,
go runtime initialize করে scheduler কে,
সেই scheduler স্কেজিউল করে করে রান করে go routine কে।
go routine রে লজিক্যাল ভার্চুয়াল thread বলা যায়।
*/
