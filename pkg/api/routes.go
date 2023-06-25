package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eluttner/api-tictac/pkg/mw"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func (s *ServerAPI) GetRoutes(ctx context.Context) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(mw.SetRequestID) // from middleare
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Please read the instructions:"))
	})

	r.Post("/game/{token}/move", s.PostGame(ctx))
	r.Get("/game", s.GetGame(ctx))
	r.Get("/game/{token}", s.GetGame(ctx))
	r.Delete("/game/{token}", s.DeleteGame(ctx))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		slog.Debug(fmt.Sprintf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares)))
		return nil
	})
	return r
}
