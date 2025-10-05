package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}

	var newProd database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProd)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Give me valid json", 400)
		return
	}
	newProd.ID = id
	database.Update(newProd)

	util.SendData(w, "Successfullny Update the value", 201)
}
