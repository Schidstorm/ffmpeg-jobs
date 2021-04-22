package server

import (
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"github.com/schidstorm/ffmpeg-jobs/api/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

type middlewareHandler struct {
	handler http.Handler
	middlewareHandlers []lib.MiddlewareHandler
}

func (m *middlewareHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	err := handlerRecursiveCall(m.middlewareHandlers, func(writer http.ResponseWriter, request *http.Request) error {
		m.handler.ServeHTTP(writer, request)
		return nil
	}, writer, request)
	if err != nil {
		logrus.Error(err)
	}
}

func handlerRecursiveCall(handlers []lib.MiddlewareHandler, next lib.MiddlewareNextHandler, writer http.ResponseWriter, request *http.Request) error {
	if len(handlers) == 1 {
		return handlers[0](writer, request, next)
	} else {
		return handlers[0](writer, request, func(writer http.ResponseWriter, request *http.Request) error {
			return handlerRecursiveCall(handlers[1:], next, writer, request)
		})
	}

}

func newMiddlewareHandler(handler http.Handler) http.Handler {
	 result := &middlewareHandler{
		handler: handler,
	}

	var handlers []lib.MiddlewareHandler
	for _, h := range middleware.Index() {
		handlers = append(handlers, h.Handle())
	}
	result.middlewareHandlers = handlers
	return result
}