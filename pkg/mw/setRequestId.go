package mw

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

const requestIDHeader = "X-Request-Id"

func SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqdId := middleware.GetReqID(ctx)
		if w.Header().Get(requestIDHeader) == "" {
			w.Header().Add(
				requestIDHeader,
				reqdId,
			)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
