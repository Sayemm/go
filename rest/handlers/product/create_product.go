package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Rquest Body")
		return
	}

	// not accessing database directly
	// dependency is injected in here, productRepo
	// using this dependency to create product
	createdProd, err := h.productRepo.Create(repo.Product{ // <<== STILL TIGHT COUPLING - Direct Access
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, http.StatusCreated, createdProd)
}
