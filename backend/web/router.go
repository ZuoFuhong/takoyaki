package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

type Router struct {
	*mux.Router
	ch *alice.Chain
}

func NewRouter() *Router {
	mw := NewMiddleware()
	chain := alice.New(mw.RequestMetricHandler)
	return &Router{
		Router: mux.NewRouter(),
		ch:     &chain,
	}
}

func (r *Router) registerHandler() {
	r.Handle("/api/columns", r.ch.ThenFunc(columnsHandle)).Methods("GET")
	r.Handle("/api/select", r.ch.ThenFunc(selectHandle)).Methods("POST")
	r.Handle("/api/insert", r.ch.ThenFunc(insertHandle)).Methods("POST")
	r.Handle("/api/update", r.ch.ThenFunc(updateHandle)).Methods("POST")
	r.Handle("/api/delete", r.ch.ThenFunc(deleteHandle)).Methods("POST")
	r.Handle("/api/datasources", r.ch.ThenFunc(ListDataSource)).Methods("GET")
	r.Handle("/api/save_datasource", r.ch.ThenFunc(AddDataSource)).Methods("POST")
	r.Handle("/api/pages", r.ch.ThenFunc(ListPage)).Methods("GET")
	r.Handle("/api/save_page", r.ch.ThenFunc(AddPage)).Methods("POST")
	// static page
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./dist"))))
}
