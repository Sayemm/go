package main

import "fmt"

func calculate() (result int) {

	fmt.Println("First: ", result) // 0

	show := func() {
		result += 10
		fmt.Println("Defer: ", result) // 15
	}
	defer show()

	result = 5
	fmt.Println("Second: ", result) // 5
	return
}

func calc() int {
	result := 0
	fmt.Println("First: ", result) // 0

	show := func() {
		result += 10
		fmt.Println("Defer: ", result) // 15
	}
	defer show()

	result = 5
	fmt.Println("Second: ", result) // 5
	return result
}

func main() {
	res1 := calculate() // 15
	fmt.Println(res1)

	res := calc() // 5
	fmt.Println(res)
}

/*
Defer
-------------------

code segment
------------
calculate := func..
CalculateAnonymous.1 := func..
calc := func...
calcAnonymous.1 := func...
main := func...

data segment
------------

stack
-----
- stack frame for main
- stack frame for calculate
	- result = 0
	- print result = 0
	- show variable = CalculateAnonymous.1 reference from code segment
	- defer show() => store this CalculateAnonymous.1
	- result = 5
	- print result = 5
	- before returning will evaluate CalculateAnonymous.1
	- stack frame for CalculateAnonymous.1
		- result value should be updated + 10
		- BUT NO RESULT VARIABLE IN THIS STACK FRAME!
		- CANNOT TAKE VALUE FROM calculate/main stack frame, only access data/code/heap
		- LAST STACK FRAME CANNNOT ACCESS PRIOR STACK FRAME. Then how result is getting updated and becoming 15?

		- defer function is stored somewhere (stack behavior but linkedList). Where?
		- ****SOMEHOW result value will be updated in calculate stack frame to 15
	- pop CalculateAnonymous.1
	- return 15
- pop calculate

- stack frame for calc
	- result = 0
	- print result = 0
	- show variable = CalcAnonymous.1 reference from code segment
	- defer show() => store this CalcAnonymous.1
	- result = 5
	- print result = 5
	- before returning will evaluate CalcAnonymous.1
	- stack frame for CalcAnonymous.1
		- result value should be updated + 10
		- now somehow without named it cannot update result value in calc or previous stack frame
	- pop CalceAnonymous.1
	- return 5
*/
