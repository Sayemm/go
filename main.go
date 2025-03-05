package main

import (
	"fmt"
)

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age = ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show // CLOSURE
}

// A closure is a function that captures variables from its surrounding scope,
// which can cause heap allocation.

// Go তে escape analysis এর মাধ্যমে function (show) এবং
// তার  parent function (outer) এর used variables (money) heap এ চলে যায়;
// তারা সবাই মিলিতভাবে closure ক্রিয়েট করে।

// Escape analysis is a way for the Go compiler to decide where to store variables
// —either on the stack or on the heap.

func call() {
	incr1 := outer()

	incr1()
	incr1()

	incr2 := outer()

	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Bank")
}

/*

1. Compilation Phase

CODE Segment
---
a
outer = func(){...}
outer:Anonymous1 = func(){...}
call = func(){...}
main = func(){...}
init = func(){...}

2. Execution Phase ->

Code Segement will be copied to RAM code segment from file code segment(./main)
-------

Data Seg
--------
p

Stack
------
stack frame create for init - print - pop from stack
stack frame for main
stack frame for call

>>>>
stack frame for outer (outer local variables will be here)
    - money, age, show (referernce to outer:Anonymous1 in code segment)
    - **money will go to heap (escape analysis) - so that we can acceess money when outer will be popped
	- **show will go to heap (escape analysis)
pop outer
incr1 will be created on call stack
	- show will be found on heap
	- then found in code through reference
stack frame for incr1 - money will be checked to local - data - code then will be found in heap
                      - money will be accessed only by incr1
					  - ** value of money will be changed
pop incr1
stack frame for incr1 - will get the new money value
pop incr1

>>>>
stack frame for outer
    - new heap for money and show after doing escape analysis
pop outer

incr2 will be created on call stack
stack frame for incr2
pop incr2
stack frame for incr2 - will get the new money value
pop incr2

pop call
pop main

Heap - GC
---------
money
show
*/
