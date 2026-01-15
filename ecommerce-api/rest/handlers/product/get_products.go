package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	if page <= 0 {
		page = 1
	}

	limitAsStr := reqQuery.Get("limit")
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	products, err := h.service.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	totalCount, err := h.service.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to count products")
		return
	}

	util.SendPage(w, products, page, limit, totalCount)
}
