package cart

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := 1

	// Clear cart
	err := h.service.ClearCart(userID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to clear cart")
		return
	}

	// Return success
	util.SendData(w, http.StatusOK, map[string]string{
		"message": "Cart cleared successfully",
	})
}
