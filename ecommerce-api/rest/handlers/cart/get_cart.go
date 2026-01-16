package cart

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := 1

	cart, err := h.service.GetCart(userID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to get cart")
		return
	}

	util.SendData(w, http.StatusOK, cart)
}
