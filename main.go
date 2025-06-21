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
	- defer show() => store this CalculateAnonymous.1 somewhere AS A CLOSURE*********
		-> (form closure show(), copy of result = 0?!!)
		-> it saves the address/pointer of result variable because parent and closure will exist in stack frame
	- result = 5
	- print result = 5
	- before returning will evaluate CalculateAnonymous.1
	- stack frame for CalculateAnonymous.1
		- copy of result = 0 + 10 = 10 ?! NO? because calcualte stack frame is still there, not popped and no value the reference of result variable is there
		- calculate exists, clousre exists and closure function is evaluating.. in this case it won't take the copy of result
		- result is the pointer of calcualte result
		- result will be 5 + 10 = 15 in calculate stack frame
	- pop CalculateAnonymous.1
	- return 15
- pop calculate

- stack frame for calc
	- result = 0
	- print result = 0
	- show variable = CalcAnonymous.1 reference from code segment
	- defer show() => store this CalcAnonymous.1 somewhere AS A CLOSURE (show(), result <-pointer)*********
	- result = 5
	- print result = 5
	==> result value will be evaluated as return (this value will be returned and stored, so does not matter if this value changes somewhere)
	- before returning will evaluate CalcAnonymous.1
	- stack frame for CalcAnonymous.1
		- result value will be updated (pointer -> 5) + 10 = 15
		- print 15
	- pop CalceAnonymous.1
	- return 5 as this was evaluted earlier and stored somewhere
- pop calc
- pop main

RULES
-----
NAMED RETURN VALUE
1. all code exexute
1. defer function will be stored on a magic box
3. return -> all defer functions will be executed
4. return the value of named return values

JUST RETURN TYPE
1. all code exexute
1. defer function will be stored on a magic box
3. return values will be evalued/stored at this time before defer (if defer even change the value does not matter)
4. all defer functions will be executed
*/
