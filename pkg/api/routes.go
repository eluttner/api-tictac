package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eluttner/api-tictac/pkg/mw"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func (s *ServerAPI) GetRoutes(ctx context.Context) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(mw.SetRequestID) // from middleare
	r.Use(middleware.Logger)

	r.Get("/", s.Home(ctx))
	r.Get("/health-check", s.HealthCheck(ctx))

	r.Post("/game/{token}/move", s.PostGame(ctx))
	r.Get("/game", s.GetGame(ctx))
	r.Get("/game/{token}", s.GetGame(ctx))
	r.Delete("/game/{token}", s.DeleteGame(ctx))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Info().Msgf(fmt.Sprintf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares)))
		return nil
	})
	return r
}
