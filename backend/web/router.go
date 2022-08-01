package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"net/url"
	"strings"
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
	r.Handle("/api/select", r.ch.ThenFunc(selectHandle)).Methods("GET")
	r.Handle("/api/insert", r.ch.ThenFunc(insertHandle)).Methods("POST")
	r.Handle("/api/update", r.ch.ThenFunc(updateHandle)).Methods("PUT")
	r.Handle("/api/delete", r.ch.ThenFunc(deleteHandle)).Methods("PUT")
	r.Handle("/api/datasources", r.ch.ThenFunc(ListDataSource)).Methods("GET")
	r.Handle("/api/datasource", r.ch.ThenFunc(AddDataSource)).Methods("POST")
	r.Handle("/api/datasource", r.ch.ThenFunc(DeleteDataSource)).Methods("DELETE")
	r.Handle("/api/datasource_option", r.ch.ThenFunc(ListDsOption)).Methods("GET")
	r.Handle("/api/pages", r.ch.ThenFunc(ListPage)).Methods("GET")
	r.Handle("/api/page", r.ch.ThenFunc(AddPage)).Methods("POST")
	r.Handle("/api/page", r.ch.ThenFunc(DeletePage)).Methods("DELETE")
	// static page
	r.PathPrefix("/").Handler(vueRouter("/", http.FileServer(http.Dir("./dist"))))
}

func vueRouter(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		// 适配Vue路由
		switch p {
		case "datasource", "page", "datatable":
			p = ""
		}
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
