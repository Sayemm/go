package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float32
}

// Pass by Value
func printing(numbers [3]int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// Pass by reference
func print(numbers *[3]int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	var address *int
	num := 20
	address = &num
	fmt.Println(address)
	fmt.Println(*address) // "dereferencing" or "indirecting"

	arr := [3]int{1, 2, 3}
	print(&arr)
	printing(arr)

	sayem_instance := User{
		Name:   "Sayem",
		Age:    24,
		Salary: 33,
	}

	p := &sayem_instance

	fmt.Println(sayem_instance.Age)
	fmt.Println((*p).Name)
	fmt.Println(p.Name) // https://go.dev/tour/moretypes/4
}
