package cart

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqAddItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (h *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	var req ReqAddItem

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.ProductID <= 0 {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	if req.Quantity <= 0 {
		util.SendError(w, http.StatusBadRequest, "Quantity must be positive")
		return
	}

	// TODO: extract user ID from JWT token
	userId := 1
	item, err := h.service.AddItem(userId, req.ProductID, req.Quantity)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.SendData(w, http.StatusCreated, item)
}
