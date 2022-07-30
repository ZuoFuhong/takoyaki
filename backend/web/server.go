package web

import (
	"log"
	"net/http"
)

type Server struct {
	name   string
	addr   string
	router *Router
}

func NewServer() *Server {
	return &Server{
		name:   "takoyaki",
		addr:   "127.0.0.1:8080",
		router: NewRouter(),
	}
}

func (s *Server) Register() {
	s.router.registerHandler()
}

func (s *Server) Serve() error {
	log.Printf("%s runs on http://%s\n", s.name, s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
