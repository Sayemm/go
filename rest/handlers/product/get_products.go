package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
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

	prductCh := make(chan []*domain.Product)

	go func() {
		productList, err := h.service.List(page, limit)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		prductCh <- productList
	}()

	ch := make(chan int64)

	go func() {
		cnt1, err := h.service.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println("Cound 1:", cnt1)
		ch <- cnt1
	}()

	totalCount := <-ch // don't need waitgroup because as long as we are not receiving published data main go routine will automatically go to sleep
	productList := <-prductCh

	util.SendPage(w, productList, page, limit, totalCount)
}

/*
Race Condition
==============
- Process (code segment, data segment, stack, heap)
- An process can have multiple threads (code, data shared)
- If all threads or goroutines try to update data from data segment at the same moment what might happen
- Might not update properly (read write)

- Race Condition
	- everyone comes at the same time and accessed shared data
	- one/multiple trited to write/update data
	- Ex: 3 thread read bank balance at the same time (5 taka), tried to deposit differnet amount, as they read the same value won't update properly

- cnt -> data segment
- multiple request (differnet person) (multiple goroutine) can come to getproducts at the same time and counter might be differnt for each request
- but all have same variable to update (global cnt) - WILL NOT GET EXPECTED RESULT
*/
