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
	Routes *chi.Mux
	Games  map[string]tictactoe.TTTIf
}

func (s *ServerAPI) ConfigureServer(ctx context.Context) {

	s.Games = make(map[string]tictactoe.TTTIf)

	s.Routes = s.GetRoutes(ctx)

}
