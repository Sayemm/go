package main

import "fmt"

func print(numbers *[3]int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	arr := [3]int{1, 2, 3}
	print(&arr)
}

/*
Code
----
print....
main....

data
----

stack
------
stack frame for array
	array of size 3 - arr

stack frame for print
    only one space for numbers, which will point/refer the first address of the array
	print array

pop print
pop main
*/
