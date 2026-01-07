package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var cnt int64

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	productList, err := h.service.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		cnt1, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println("Cound 1:", cnt1)
		cnt = cnt1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cnt2, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println("Cound 2:", cnt2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cnt3, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println("Cound 3:", cnt3)
	}()

	wg.Wait()

	util.SendPage(w, productList, page, limit, cnt)

	go hei(wg) // DON'T COPY WaitGroup
}

/*
Waitgroup
=========
wg.Add(1) -> wg counter (state) incrase 1
wg.Done() == wg.Add(-1)

- main go routine -> 3 go routine
- wg.Wait()
	-> checks the value of waitgroup
	-> if != 0 => MAIN GO ROUTINE WILL GO TO SLEEP
	-> when wg will become 0 (all 3 go routines will be done)
		- wg was created on the main go routine
		- so main go routine will wake up from sleep
		- go runtime will do that (awake main go routine)

- if we forget to add/remove to the waitgroup, application will be run infinite or panic/crash
- multiple defer? - execute the last defer first

WaitGroup Internals
===================
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}

- wg instance, this struct has 3 receiver function
- wg will get a default value as we didn't use *sync.Waitgroup
- var wg sync.WaitGroup ===>
- wg = WaitGroupP{
			noCopy: noCopy{},
			state: atomic.Uint64{
				_: noCopy{},
				_: align64{},
				v: 0
			},
			sema: 0
		}

- wg.Add(3)
	- v has 64 bits
	- 32 bit(high) | 32 bit(low)
	- high 32 bit = counter / unfinished worker counter / unfinished go routines
	- 3 => 000..11 | 32 bit

- wg.Wait()
	- it will check the value on high 32 bit
	- if 0 (high 32), main go routine will not sleep
	=> if !0 - main go routine will go to sleep (HOW?)
		-> runtime_SemacquireWaitGroup(&wg.sema)
		-> CALLS GO RUNTIME AND SENDS THE ADDRESS OF wg.sema (A5-assume) and TELLS TO (ACQUIRE) SLEEP THE MAIN GO ROUTINE (from where the wg.Wait() is called)
		-> ADDRESS is UNIQUE (wg.sema variable address is Unique, value 0 - we are just using the address)
		-> runtine keeps a global map (map[uint32][]Goroutine) - map of list
			-> mp[A5] = mp[A5].append(mp[A5], Goroutine{})
			-> low 32 bit of v will increase 1 (keeps the cunter for how many go routines called wg.wait())

- wg.Done()
	- wg.Add(-1)
	- high 32 of v will decrease
	- if high 32 of v become 0
		-> for ; w != 0; w-- {
				runtime_Semrelease(&wg.sema, false, 0)
			}
		-> w = low 32 bit value (low 32 keeps the sleeping goroutine counter)
		-> loop through all the sleep mode goroutines
		-> runtime_Semrelease(&wg.sema, false, 0)
			-> telling to release those goroutines (awake)

- noCopy -> tells not to copy the waitgroup (warning)
	- copy of wg, another struct another memory address
	- address of sema will also change
	- might create deadlock
	- then why not error?
		-> first time (no add/done yet) copy okay (high 0 and low 0 - no problem)
*/

func hei(wg sync.WaitGroup) {
	wg.Wait()
}

/*
WaitGroup main goroutine-এর stack frame-এ initialize হয়।
main goroutine যদি Wait() কল করে sleep/block হয়ে থাকে,
তাহলে Wait() কীভাবে “constantly” counter check করে?
কে জানে কখন Done() হলো?

----
wg নামের ভ্যারিয়েবলটা main goroutine-এর stack-এ থাকে
কিন্তু sync.WaitGroup–এর ভিতরের state (counter + semaphore) stack-এ থাকে না
closure + escape analysis এর কারণে সেটা heap-এ চলে যায়, তাই সব goroutine শেয়ার করতে পারে

wg কে goroutine capture করছে (closure)
	তাই wg stack-এ রাখা unsafe
	compiler wg কে heap-এ move করে
এখন:
	main goroutine
	anonymous goroutine
	দুজনেই same heap-allocated WaitGroup ব্যবহার করছে
----


WaitGroup-এর ক্ষেত্রে wg.Wait() কোনোভাবেই counter-কে বারবার check করে না বা busy-loop চালায় না।
যখন main goroutine Wait() কল করে, তখন সে Go runtime-এর একটি semaphore-এর উপর block (park) হয়ে যায়
এবং scheduler তাকে CPU থেকে সরিয়ে দেয়।

WaitGroup-এর counter (state) heap-এ থাকা একটি shared atomic variable, যেটা সব goroutine access করতে পারে।
প্রতিবার কোনো worker goroutine wg.Done() কল করলে, সে atomicভাবে counter কমায়,
আর যে goroutine-টি শেষবার Done() কল করে counter-কে 0 বানায়, সেই goroutine-ই runtime-কে signal দেয় (semaphore release করে)।
তখন runtime scheduler আগে থেকে sleep/park করে রাখা main goroutine-কে আবার ready করে তোলে।
ফলে main goroutine নিজে কিছু check না করেও, runtime + atomic state + semaphore ব্যবস্থার মাধ্যমে জানতে পারে যে সব goroutine শেষ হয়ে গেছে।

*/
