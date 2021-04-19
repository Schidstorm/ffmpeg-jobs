package server

import (
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"net/http"
)

type Server struct {
	handler *HttpHandler
}

func NewServer() *Server {
	return &Server{
		handler: NewHttpHandler(),
	}
}

func (s *Server) Serve(addr string) error {
	server := &http.Server{Addr: addr, Handler: s.handler.Handler()}
	return server.ListenAndServe()
}

func (s *Server) AddController(controller lib.Controller) {
	s.handler.AddController(controller)
}
