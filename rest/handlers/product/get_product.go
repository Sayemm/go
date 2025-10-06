package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	product, err := h.productRepo.Get(id)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	util.SendData(w, http.StatusOK, product)
}
