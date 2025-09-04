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

// why Manager? we can just call With? like middleware.With?
// The Manager has a globalMiddlwares slice, so we could register middlewares that are automatically applied to every route,
// Instead of passing middlewares inline every time, you can let the Manager keep track of defaults and just add route-specific ones when needed.

func (mngr *Manager) With(middlewares ...MiddleWare) MiddleWare { // When we will call this MiddleWare list become closure and create a middlewares variable in that closure
	return func(next http.Handler) http.Handler {
		n := next
		for i := len(middlewares) - 1; i >= 0; i-- { // [hudai, logger]
			middleware := middlewares[i]
			n = middleware(n)

			// 1: n = middleware.Logger(n = http.HandlerFunc(handlers.Test))
			// 0: n = middleware.Hudai(n = middleware.Logger(http.HandlerFunc(handlers.Test)))
		}
		return n
	}
}
