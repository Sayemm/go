package product

import (
	"ecommerce/domain"
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

	if req.Title == "" {
		util.SendError(w, http.StatusBadRequest, "Title is required")
		return
	}
	if req.Price < 0 {
		util.SendError(w, http.StatusBadRequest, "Price must be positive")
		return
	}

	product := domain.Product{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	}

	updatedProduct, err := h.service.Update(product)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	if updatedProduct == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, http.StatusOK, updatedProduct)
}
