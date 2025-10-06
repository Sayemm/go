package main

import (
	"fmt"
	"os"
)

// import (
// 	"ecommerce/cmd"
// )

// Only Signature of the function
type People interface {
	PrintDetails()
	ReceiveMoney(money float64) float64
}

type BankUser interface {
	WithDraw(amount float64) float64
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

func (obj user) WithDraw(amount float64) float64 {
	return obj.Money - amount
}

func main() {
	var u1 People // PrintDetails function must be present as a reciever function of that struct
	u1 = user{    // instantiation - object/instance creation
		Name:  "Sayem",
		Age:   26,
		Money: 23,
	}
	u1.PrintDetails()
	fmt.Println(u1.ReceiveMoney(64))

	var u2 BankUser
	u2 = user{
		Name:  "Sayem",
		Age:   26,
		Money: 23,
	}
	fmt.Println(u2.WithDraw(2))

	// u2.PrintDetails() XXX - CANNOT DO THIS
	// cannot use u2 as a People as it is BankUser interface
	// we can do it by converting u2 to a struct object
	obj, ok := u2.(user) // is u2 object of user struct
	if !ok {
		fmt.Println("User 2 is not type of user struct")
		os.Exit(1)
	}
	obj.PrintDetails()
	fmt.Println(obj.ReceiveMoney(4))
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
- Interface is the parent of struct
- Struct implements/extends the interface
*/
