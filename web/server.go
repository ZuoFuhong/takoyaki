package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	traceKey = "traceId"
)

type Server struct {
	addr      string
	threshold int
}

type middlewareHandler struct {
	m *http.ServeMux
}

func NewMiddlewareHandler(mux *http.ServeMux, cc int) http.Handler {
	middle := middlewareHandler{
		m: mux,
	}
	return middle
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 注入traceId
	traceId, _ := uuid.NewUUID()
	r = r.WithContext(context.WithValue(context.Background(), traceKey, traceId.String()))
	m.m.ServeHTTP(w, r)
}

func NewServer(addr string, threshold int) *Server {
	return &Server{
		addr:      addr,
		threshold: threshold,
	}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./frontend/dist"))))
	mux.HandleFunc("/api/columns", columnsHandle)
	mux.HandleFunc("/api/select", selectHandle)
	mux.HandleFunc("/api/insert", insertHandle)
	mux.HandleFunc("/api/update", updateHandle)
	mux.HandleFunc("/api/delete", deleteHandle)
	middle := NewMiddlewareHandler(mux, s.threshold)
	err := http.ListenAndServe(s.addr, middle)
	return err
}
