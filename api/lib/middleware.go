package lib

import "net/http"

type MiddlewareNextHandler = func(http.ResponseWriter, *http.Request) error
type MiddlewareHandler = func(http.ResponseWriter, *http.Request, MiddlewareNextHandler) error

type Middleware interface {
	Handle() MiddlewareHandler
}
