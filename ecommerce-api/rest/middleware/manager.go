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

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(mux http.Handler) http.Handler {
	n := mux

	for _, middleware := range mngr.globalMiddlewares {
		n = middleware(n)
	}

	return n
}
