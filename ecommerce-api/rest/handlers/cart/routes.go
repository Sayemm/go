package cart

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /cart/items", manager.With(
		http.HandlerFunc(h.AddItem),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("GET /cart", manager.With(
		http.HandlerFunc(h.GetCart),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("PUT /cart/items/{id}", manager.With(
		http.HandlerFunc(h.UpdateItem),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("DELETE /cart/items/{id}", manager.With(
		http.HandlerFunc(h.RemoveItem),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("DELETE /cart", manager.With(
		http.HandlerFunc(h.ClearCart),
		h.middlewares.AuthenticateJWT,
	))
}
