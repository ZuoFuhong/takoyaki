package web

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"takoyaki/defs"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) RequestMetricHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Inject traceId to context
		traceId, _ := uuid.NewUUID()
		r = r.WithContext(context.WithValue(context.Background(), defs.TraceKey, traceId.String()))
		log.Printf("%s %q\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
