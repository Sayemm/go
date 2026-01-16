package cart

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	// Get item ID from URL
	itemIDStr := r.PathValue("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	// Remove item
	err = h.service.RemoveItem(itemID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Return success
	util.SendData(w, http.StatusOK, map[string]string{
		"message": "Item removed successfully",
	})
}
