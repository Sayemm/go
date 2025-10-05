package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}

	database.Delete(id)

	util.SendData(w, "Successfullny Deleted the Product", 201)
}
