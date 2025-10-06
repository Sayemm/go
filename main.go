package main

import "fmt"

// import (
// 	"ecommerce/cmd"
// )

// Only Signature of the function
type People interface {
	PrintDetails()
	ReceiveMoney(money float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

// receiver function / method / behavior
func (obj user) PrintDetails() {
	fmt.Println(obj.Name)
}

func (obj user) ReceiveMoney(money float64) float64 {
	return obj.Money + money
}

func main() {
	var u1 People // PrintDetails function must be present as a reciever function of that struct

	u1 = user{ // instantiation - object/instance creation
		Name:  "Sayem",
		Age:   26,
		Money: 23,
	}
	u1.PrintDetails()
	fmt.Println(u1.ReceiveMoney(64))
}

/*
Singleton Design Pattern
------------------------
- everyone is sharing ONE thing/object
- We are loding config only one time using pointer


Paradigms
---------
- Structured Programming Language
- Object Oriented Programming Language
- Functional Programming Language
	- Pure Functional Paradigm - same input same output always

Interface
=========
- Abstractive - Concept (dharona)
- Interface - Pure Abstraction, No details

Struct
======
- struct can implement interface
- struct and interface er type same hoe jay if the struct has a same receiver function
*/
