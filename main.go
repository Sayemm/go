package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUser(usr User) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

// RECEIVER FUNCTION - Works only with custom type
// only User type variable can call this function - user1.printDetails()
func (usr User) printDetails() {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Sayem",
		Age:  30,
	}

	printUser(user1)
	user1.printDetails()
	user1.call(5)

	user2 := User{
		Name: "Roki",
		Age:  60,
	}

	// printUser(user2)
	user2.printDetails()

}

/*
Code Seg
--------
User
printUser()
printDetails() //User
call() //User
main()

Data Seg
--------

Stack
-----
main stack frame
  - user1

printUser stack frame
  - usr (copy of user1)
pop printUser
printDetails stack frame
  - usr (copy of user1)
pop printDetails
call stack frame
  - usr (copy of user1)
  - a (copy)
pop call

main stack frame
  - user2

printDetails stack frame
  - usr (copy of user2)
pop printDetails

pop main

Heap
----
*/
