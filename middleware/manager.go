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

func (mngr *Manager) With(next http.Handler, middlewares ...MiddleWare) http.Handler {
	n := next
	// [logger, hudai]
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n

	// n = http.HandlerFunc(handlers.Test)
	// n = middleware.Logger(http.HandlerFunc(handlers.Test))
	// n = middleware.Hudai(n = middleware.Logger(http.HandlerFunc(handlers.Test)))
}
