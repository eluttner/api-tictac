package api

import (
	"context"

	"github.com/eluttner/api-tictac/pkg/tictactoe"
	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	//here we can add more config options like databases, queues, 3rd part services
}

type ServerAPI struct {
	Config ServerConfig
	Games  map[string]tictactoe.TTT
}

func (s *ServerAPI) ConfigureServer(ctx context.Context) *chi.Mux {

	s.Games = make(map[string]tictactoe.TTT)

	r := s.GetRoutes(ctx)
	return r
}
