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

	p := func(a int) {
		fmt.Println(a)
	}
	defer p(result)

	defer fmt.Println(result)       // 5
	fmt.Println("Second: ", result) // 5
	defer fmt.Println(5)

	return
}

func main() {
	res1 := calculate() // 15
	fmt.Println(res1)
}

/*
Defer
-------------------

code segment
------------
calculate := func..
CalculateAnonymous.1 := func..
CalculateAnonymous.2 := func..
main := func...

data segment
------------

stack
-----
- stack frame for main
- stack frame for calculate
	- result = 0
	- print result = 0
	- show = CalculateAnonymous.1 reference from code segment
	- another cell for defer => defer list pointer => some address
	* defer show() => store this CalculateAnonymous.1, result somewhere (maybe HEAP) AS A CLOSURE*********
	- result = 5
	- p = CalculateAnonymous.2 reference from code segment
	* CalculateAnonymous.2 will be stored in heap with argument as result = 5
	* Println will be storead in heap with argument = 5
	- print second 5
	* Println will be storead in heap with argument = 5

	- defer functions will be exeuted using defer list pointer..
	- and defer list pointer will be updated from the next pointer until it becomes nil
- pop calculate
*/
