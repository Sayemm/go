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

/* RO RUNTIME - mini OS
- RAM (Kernel Space, User Space)
	- Kernerl Space
		-> OS exists there
		-> User spcae cannot access kernel space
	- User Space
		-> other application
- if any application that is running in user space, needs a file from HD
	-> it cannot access it directly
	-> It requests kernel for that file
	-> This request is called System call

	-> Ram/Process ask (read) for that file to kernel
	-> OS puts that data in a buffer (user space)
	-> This data has a file descriptor (number > 3)
	-> Kernel gives this fd to the process
	-> Process then request using fd
	-> Kernel then gives data from that buffer to that process for reading

- if process from user space, needs anything like file, socket, network related anyting.. it requests that to the kernel

-------------

- ./main - go process will be created
- this process will allocate some RAM (code, data segment) in user space
- OS will start to run this process / MAIN THREAD
	-> Process starts means - Main THREAD will start
	-> MAIN THREAD STARTS - PROCESS STARTS

- When main thread will execute a code will execute
	-> this code is not the code written in main.go
- GO RUNTIME CODE WILL EXECUTE FIRST
- main thread executes go runtime code
	-> stack, heap creates...stack frame...
- GO RUNTIME DOES SOME WORK
	-> GO Scheduler initialize
	-> System call to Kernel for epoll_create
		-> epoll is a feature of Kernel
		-> epoll has 3 operation (epoll_create, epoll_ctl, epoll_wait)
		-> epoll_create means go runtime tell kernal to create separate OS thread (now 2 thread - MAIN and another one)
		-> this 2nd thread will just do epoll_wait that is always in sleep mode
			-> meaning if main thread do epoll_ctl to the kernal and when fd is ready it will awake the 2nd thread
			-> epoll_wait send this fd to go runtime
			-> now go runtime/main thread request to read that file using that fd
		-> go runtime has a limitation like 100.. ie. it can request to read 100 file parallely to the kernal
		-> only kernal can control this thread, go runtime does not have control over this 2nd thread
	-> setup GC
		-> another 3rd separate OS thread will be created
		-> that thread will run GC
		-> go runtime has control over this GC thread
	** after doing everying go runtime will continue to do scheduling



- VCPU runs thread
- Let's say total 6 thread (T1...T6)
- T1 wants to read a file, then it will sys call to the Kernal
- kernal will return FD to that thread
- KERNAL does not do it right away, it might take some time to give that fd to that process
- NOW epoll/kqueue/IOCP comes into the scene
	- thread requests epoll_ctl to the kernal
	- epoll send that thread to sleep mode
- VCPU will now run another thread
- after doing everything Kernal will awake that sleeping thread and send the FD to epoll_wait (epoll_wait will be in that sleeping thread)
- now thread has fd, now it will read request to Kernal and ultimately will be able to read that file

GO SCHEDULING
-------------
- go runtime maintains a rule - m:p:g
- m = machine/os thread
- p = logical/virtual processor (go create this - has v irtualalu, cu...)
- g = go routines
-> 4 vcpu = m = 4 and p will be 4 (go will create 4 logical processor)
-> let's say 16 go routines
- 4:4:16
- os thread (4) runs virtual processor (4) and virtual processor is running 16 goroutines

- 2 VCPU - 2 more OS thread will be created in GO process for scheduling
- 2 virtual will also be created and go runtime only knows about that
- 2 vcpu will run 2 virtual/logical processor separately
- each p has a queue -> local run queue
		- run queue has a slot of 256
		- run queue is circular
		- go routine will come into this queue
		- logical processor (p) will give priority to the first go routine then next... (run concurrently - pcb - push to the queue again..)
- there is a global run queue that is dynamic
		- as long as it has memory we can push go routine here
		- when go routine is being creted, it takes place on the available local run queue
		- when local run queue is full it will go to global run queue
		- when there is no go routine for a logical processor it takes go routine from another logical processors local run queue
		- if logical processor does not have anything to do (on his own local run queue and others' run queue) then it will go to the global run queue


=========================
- ./main
- main go routune will be created that will run go runtime
- now go runtime will create another thread (epoll_wait) and another thread for GC
- go scheduler maintains m:p:g and creates m OS Thread that will run go virtual processor

- go runtime now will run all init function
- then creates another goroutine (main [main.go] goroutine) that runs on heap (on the stack main thread is running go runtime)
	- main go routine will go to the local run queue of one of the virtual processor
	- stack for that main go routine and main stack frame create
		- mux variable on main function
		- NewServeMux stack frame
		- serve stack frame (there is an infinite loop in there)
			- Accept stack frame
				- goroutine tells go runtime to create a socket
				- go runtime tells (sys call - epoll_ctl/netpoll) Kernal to create a socket, before go runtime will also check whether the socket is already created or not.
				- go runtimes just tell the Kernal to create the socket and leave (go runtime keeps doing it's work)
				- go runtime then sends that go routine to SLEEP
				- socket is line a pipe (both way)
				- Kernal creates a FD and corresponding socket and give it an endpoint (8080)

				- go routine AWAKE - request to that URL (localhost:8080)
				- router receives the request then request goes to computer/servers NIC
				- NIC writes this in the Receice Buffer (on the RAM) then interrupt request to kernal
				- Kernal reads from NIC receive bufffer and writes this in that socket receive buffer AND marks the FD of that socket
				- then Kernal awake the 2nd thread (epoll_wait) and send that FD - epoll_wait thread send that FD to go runtime
				- go runtime finds out who are sleeping for that particular FD and awake that go routine and give that FD
				- now that sleeping goroutine started to run again by a virtual processor
				- now that go routine reads DATA from the socket buffer that is mapped by that FD (rw, err := l.Accept()) and keep it on rw
				- now that man go routine created another goroutine (go c.serve(connCtx)) and give it the responsibility
					- new stack will be allocated for new go routine
					- new goroutine will find the path /about and finds out which handlers/route are being registered in the router and match
					- and write data to socket send buffer
					- Kernal writes this data to NIC ring buffer
					- NIC send data to router and finally sender will see the data
				- main go routine will be ready again and call l.Accept().., tells to socket again but go runtime will find that socket is already created
				- it will just wait for kernal to mark the socket buffer



*/
