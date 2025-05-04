package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	//x = 1, 2, 3, 4 (len4, cap4)
	//y = 1, 2, 3 (len3, cap4)
	y = append(y, 5)
	//x = 1, 2, 3, 5 (len4, cap4)
	//y = 1, 2, 3, 5 (len4, cap4)

	x[0] = 10
	//x = 10, 2, 3, 5 (len4, cap4)
	//y = 10, 2, 3, 5 (len4, cap4)

	fmt.Println(x)
	fmt.Println(y)
}

/*
main stack frame = x[ptr = nil, len = 0, cap = 0]
---->
append stack frame = 1
		 s[ptr = nil, len = 0, cap = 0] <-- copy x
         [1] -> escape analysis hoe heap e jabe

		 s1[ptr = heap address, len = 1, cap = 1]
		 s1 return and goes to x
pop append stack fram => heap = [1]

x[ptr = heap address, len = 1, cap = 1]


---->
append stack frame = 2
		 s[ptr = heap address, len = 1, cap = 1] <-- copy x
         allocate 1 x 2 = capacity - heap [1] [_1 _2] => new array create and copy from previous

		 s1[ptr = heap address_1, len = 2, cap = 2]
		 s1 return and goes to x
pop append stack fram => heap = [1, 2]

x[ptr = heap address_1, len = 2, cap = 2]


---->
append stack frame = 3
		 s[ptr = heap address_1, len = 2, cap = 2] <-- copy x
         allocate 2 x 2 = capacity - heap [1] [_1 _2] [*1 *2 *3 *]=> new array create and copy from previous

		 s1[ptr = heap address*1, len = 3, cap = 4]
		 s1 return and goes to x
pop append stack frame => heap = [1, 2, 3]

x[ptr = heap address*1, len = 3, cap = 4]


---> y := x
x[ptr = heap address*1, len = 3, cap = 4]
y[ptr = heap address*1, len = 3, cap = 4]

--->
append stack frame = 4
		 s[ptr = heap address*1, len = 3, cap = 4] <-- copy x
         [*1 *2 *3 *4]

		 s1[ptr = heap address*1, len = 4, cap = 4]
		 s1 return and goes to x
pop append stack fram => heap = [1, 2, 3, 4]

x[ptr = heap address*1, len = 4, cap = 4]


--->
append stack frame = 5
		 s[ptr = heap address*1, len = 3, cap = 4] <--copy y
         [*1 *2 *3 *5] => 5 will append after len 3 in 4th position

		 s1[ptr = heap address*1, len = 4, cap = 4]
		 s1 return and goes to y
pop append stack fram => heap = [1, 2, 3,5]
y[ptr = heap address*1, len = 4, cap = 4]

---> x[0] = 10

x[ptr = heap address*1, len = 4, cap = 4]
heap = [10, 2, 3, 5] => x pointing to 1st element of heap address that will change


==> slice underlying array rule => keep doubling (100% till 1024 cap then 25%)
1024 cap = 1024 x 2 = 2048
2048 cap = 2048 * 1.25 = 2560
*/
