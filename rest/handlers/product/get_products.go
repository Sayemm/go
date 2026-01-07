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


*/
