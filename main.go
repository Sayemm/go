package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is Sayemmm")
}

func main() {
	mux := http.NewServeMux() // mux = router

	mux.HandleFunc("/hello", helloHandler) // /hello - route
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on:3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}

}

/*
- We can create multiple route using router
- request goes to router and router knows which request will go to which route
- then the function that is registered with that router will be executed

- Core of OS is kernel
- When we request to server from a client
	- router is attached to that server
	- request goes to router
	- router send the request to the server (wifi adapter - NIC)
	- Kernel handles NIC
	- NIC translate it to binary data
	- Kernel keeps some memory in RAM for NIC
		- name of this memory is Write Buffer
	- NIC keeps that binary data to Write Buffer includig some metadata like sender ip address, port etc.

	- NIC request an interrupt signal to Kernel
	- Kernel understands now that a request has come to NIC and NIC already saved it to RAM
	- Kernel will read and copy the data from Write buffer
	-
		- In RAM some spaces are allocated for socket
		- When go server runs, it creates a socket immediately ->
		- and that socket has buffer (socket receive buffer)
	-
	- data that was copied from write buffer by kernel will be be written in socket receive buffer (3000port)
	-
		-> ListenAndServe -- l.Accept() -> tell go runtime that create a socket for me
		- go runtime creates a socket by requesting Kernel to create that socket
		- Kernel creates a socket and store it to open fine descriptor (8, socket..)
		- Kernel give fd = 8 to go runtime
		- socket has endpoint port 3000
		- now l.Accept() funciton goes to sleep mode

		- request comes with lots of information
		- Kernel decides the port - 3000, which socket port to take that write buffer data (port number is mentined in write buffer data)
		- A computer can have lots of process, and every process can have multiple socket
		- Kernel finds out the socket for 3000 endpoint, that is mentioned in write buffer
		- Now kernel saves that data to 3000 port socket's receive buffer
		- Now Kernal makes that socket file decriptor readable
		- Kernel mentiones this to go runtime

	- Kernel makes the file readable in it's open file table for that socket file (fd = 8)
	- Means there is a request in the socket
	- This information is now sent to go runtime by kernel that fd = 8 is not readable
	- now go runtime finds out which go routine requested for number 8 socket
	- now go runtime tells main go routine to wake up from sleep. l.Accept()
		- main go routine reads data from that socket receive buffer (as fd = 8 is now readable)
		- l.Accept() reads the data after waking up
		-
	- rw gets that data, main go routine created another go routine
	- now that data is being handled by new go routine
		- now new go routine looks for router (mux)
		- finds the registered route (/about, /hello)
		- router matches the data that came fron receive buffer
		- new go routine executes that handler function (helloHandler/aboutHandler)

		- Socket has send buffer as well.
		- Socket interrupts kernel that now I have data ("Hello world")
		- Kernel copies send buffer and send it to NIC
		- NIC has ring buffer and that data is saved to ring buffer
		- rind buffer is file - has fd - now that is readable
		- NIC reads data from ring buffer
		- NIC sends this data (01010) to router
		- NIC send it as a response to the client

	- now again re, err := l.Accept() - socket already created
	- tell go runtime that I am waiting to number 8 and going to sleep again.

File Descriptor (number > 0) - Who can identify file uniquely
--------------------------------------------------------------
- OS has kernel
- Process can have certain size of file descriptor
- default is 1024, process can have 1024 file descriptor
- File: where we store data in HD/RAM

When go prcess wants to open a file from HD
	- send request to kernel to access that file
	- Kernel has open file table (column: fd, file) (7, a.txt)
	- Kernel will give number 7 to that process, now process has 7

	- Now process can say to kernal that I want to read fd-7
	- Each process has there own FD

SOCKET
------
- in Linux socket itself is a file
- it's like a pipe which can send or receivse data (process ---> <--- pipe --> <---)
- socket has protocol (tcp, udp)
- process can create socket to cmmunicate with the client
- process will tell kernel to create a socket
- kernel will create a socket - kernel will add it to open file table - (fd, file) (8, socket file) and send fd = 8 to process
- when any request comes and if that process wants to read that request information - then process has to tell kernel that I want to do some operation with the file with fd = 8
- if kernel permits that process will be able to do that operation

*/
