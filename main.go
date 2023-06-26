package main

import (
	"context"
	"net/http"
	"os"

	"github.com/eluttner/api-tictac/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Configure web server
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	ctx = logger.WithContext(ctx)

	s := &api.ServerAPI{}
	s.ConfigureServer(ctx)

	log.Info().Msg("Starting server")
	err := http.ListenAndServe(":3000", s.Routes)
	if err != nil {
		log.Error().Err(err).Msg("Error: starting the server")
	}

	//No need for graceful shutdown, not using any database or queue
}
