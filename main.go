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
	-> setup GC
		-> another 3rd separate OS thread will be created
		-> that thread will run GC
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

*/
