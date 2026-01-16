package cart

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateItem struct {
	Quantity int `json:"quantity"`
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get item ID from URL
	itemIDStr := r.PathValue("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	// Parse request
	var req ReqUpdateItem
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Validate quantity
	if req.Quantity <= 0 {
		util.SendError(w, http.StatusBadRequest, "Quantity must be positive")
		return
	}

	// Update item
	err = h.service.UpdateItem(itemID, req.Quantity)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Return success
	util.SendData(w, http.StatusOK, map[string]string{
		"message": "Item updated successfully",
	})
}
