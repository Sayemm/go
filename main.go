package main

import "fmt"

type User struct { // readonly - not possible to change user type later - code seg
	Name string // member variable / property
	Age  int
}

func main() {
	var user1 User

	user1 = User{ // user2 is instance or object of User type
		Name: "Sayem",
		Age:  30,
	}

	fmt.Println(user1.Name)

	user2 := User{ // instantiate: process of creating instance
		Name: "Roki",
		Age:  60,
	}

	fmt.Println(user2.Age)
}

/*
Code Segment - readonly
------------
User = type User struct {....}
main = func() {....}

Data Segment
------------
stack frame for main
    - user1
	- user2
pop main
*/
