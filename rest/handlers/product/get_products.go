package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

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

	// Takes 4 Seconds
	cnt, err := h.service.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	fmt.Println("Cound 1:", cnt)

	go func() {
		// Takes 4 Seconds
		cnt1, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(cnt1)
	}()

	go func() {
		// Takes 4 Seconds
		cnt2, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(cnt2)
	}()

	time.Sleep(4 * time.Second)

	/*
		Sequentially would take 12 seconds (4*3)
		Concurrenlty will take 8 seconds (4 (without goroutine) + 4 (2 go routine))
	*/

	util.SendPage(w, productList, page, limit, cnt)
}

/*
Why goroutine and go channel matters
------------------------------------
- go rutime -> main go routine - stays on the STACK
- HEAP -> other go routines (create stack frame for functions and other stuff...)

      1 Req takes 1.34kB in the RAM
   1000 Req tak (1000*1.34) ~ 1Mb
  10000 Req ~ 10Mb
 100000 Req ~ 100Mb
1000000 Req ~ 1.34GB

- So the server where we are running the go process
the RAM in that device has to be more that 1.34GB

- So we need to find out how much RAM a process can take
- We also need to find out whether database can handle that much request or not
	- how much data database is loading to the RAM
- How server handles the request?
	- CPU handles the request
	- more VCPU can handle more requests

- lets' say server cannot handle 1M request in a second, (lets say it takes 10second to handle 1M request)
  what would happens if I keep sending 1M request in every second
	- 1s hold 1Gb data, 2s 2GB.....10s 10GB
	- handled request will freed the ram but new request will occopy the RAM again
	- if i don't have 10GB RAM, Server will CRASH!
*/
