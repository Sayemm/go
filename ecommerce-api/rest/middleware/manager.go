package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler // signature type Sayem int // type Sayem func(a int)
type Manager struct {
	globalMiddlewares []MiddleWare
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]MiddleWare, 0),
	}
}

func (mngr *Manager) Use(middlewares ...MiddleWare) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...MiddleWare) http.Handler {
	n := next

	// Apply middleware in reverse order
	// So they execute in the order you specified
	for i := len(middlewares) - 1; i >= 0; i-- {
		n = middlewares[i](n)
	}

	return n
}

func (mngr *Manager) WrapMux(mux http.Handler) http.Handler {
	n := mux

	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
		n = mngr.globalMiddlewares[i](n)
	}

	return n
}
