package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	x := []int{1, 2, 3, 4, 5} // [1 2 3 4 5] len5 cap5

	x = append(x, 6) // [1 2 3 4 5 6] len6 cap10
	x = append(x, 7) // [1 2 3 4 5 6 7] len7 cap10

	a := x[4:] // [5 6 7] len3 cap6 (because x has cap10)

	y := changeSlice(a)

	fmt.Println(x) // [1 2 3 4 10 6 7] len of x is 7
	fmt.Println(y) // [10 6 7 11] len of y is 4 after the append
	fmt.Println(a) // [10 6 7] len of a is 3

	fmt.Println(x[:8]) // [1 2 3 4 10 6 7 11]

	// len change when append happens
}

/*
main stack frame
= array create [1, 2, 3 4, 5]
= x[ptr = address of 1, len = 5, cap = 5]

---->
append stack frame = 6
		s[ptr = address of 1, len = 5, cap = 5] <-- copy x
        allocate 5 x 2 = 10 capacity array heap e create

		s1[ptr = heap address 1, len = 6, cap = 10]
		s1 return and goes to x
pop append stack fram => heap = [1 2 3 4 5 6]

x[ptr = heap address 1, len = 6, cap = 10]


---->
append stack frame = 7
		s[ptr = heap address 1, len = 6, cap = 10] <-- copy x

		s1[ptr = heap address 1, len = 7, cap = 10]
		s1 return and goes to x
pop append stack fram => heap = [1 2 3 4 5 6 7]

x[ptr = heap address 1, len = 7, cap = 10]


---->
a[ptr = heap address 5, len = 3, cap = 6]

---> y := changeSlice(a)
stack frmae for changeSlice
p[ptr = heap address 5, len = 3, cap = 6]
heap address 5 = value will change from 5 to 10 in this address
head = [1 2 3 4 10 6 7]

append 11
---------
s[heap address 10, len = 3, cap = 6]
s1[heap address 10, len = 4, cap = 6]
return s1 to y
y[heap address 10, len = 4, cap = 6]
head = [1 2 3 4 10 6 7 11]

pop append
pop changeSlice
*/
