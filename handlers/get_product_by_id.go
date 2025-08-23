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

	for _, product := range database.ProductList {
		if id == product.ID {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "No Data", 404)
}
