package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	_, err = h.productRepo.Update(repo.Product{ // <<== STILL TIGHT COUPLING - Direct Access
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfullny Update the value")
}
