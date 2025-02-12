package main

import "fmt"

/* Higher Order Function / First Class function
1. Parameter -> Function or
2. Return function or
3. both
*/

// Callback Function: Function that we pass on Higher Order function as parameter or argument

func processOperation(a int, b int, op func(x, y int)) { // Callback Function: op func(x, y int)
	op(a, b)
}

func add(x, y int) {
	fmt.Println(x + y)
}

func call() func(a, b int) {
	return add
}

func main() {
	processOperation(5, 6, add)

	rt := call()
	rt(3, 4)
}

/*
Math (Discrete Math) -> Functional Paradigm -> go
       |
1. First Order Logic
2. Higher Order Logic

First Order Function
	1. Standard Function or Named Function
	2. Anonymous Function
	3. IIFE
	4. Function Expression

Logic
1. Object (Car, People)
2. Property (Color, Student)
3. Relation

### First Order Logic -> works with Object, Property and Relation (just Rule)
Rule: All customer must pay the bill
      All student must wear the uniform


### Higher Order Logic -> Works with Rules as well
   Any rule that applies to all cx must also apply to Tutul
   Rule: All cx must pay tips to the waiters -> Tutul will also pay the tips

First class Citizen -> things that we can assign to a variable
*/
