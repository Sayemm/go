package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}

	product := database.Get(id)
	if product == nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}

	util.SendData(w, product, 200)
}
