package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/eluttner/api-tictac/pkg/tictactoe"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rs/zerolog/log"
)

func (s *ServerAPI) GetGame(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sublogger := log.With().
			Str("reqId", middleware.GetReqID(r.Context())).
			Logger()
		token := chi.URLParam(r, "token")
		if token != "" {
			sublogger.Info().Msgf("Token: %s", token)
		}

		if token == "" {

			var mutex = &sync.RWMutex{}
			mutex.Lock()
			t := &tictactoe.TTT{}
			token := t.NewGame()
			s.Games[token] = t
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t.GetGame())
			return
		} else {
			var mutex = &sync.RWMutex{}
			mutex.RLock()
			if val, ok := s.Games[token]; ok {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(val.GetGame())
				mutex.RUnlock()
				return
			}
			mutex.RUnlock()
		}
		w.WriteHeader(http.StatusNotFound)
	})
}

func (s *ServerAPI) PostGame(ctx context.Context) http.HandlerFunc {
	type Req struct {
		Token  string `json:"token"`
		Player string `json:"player"`
		Row    int    `json:"row"`
		Column int    `json:"column"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sublogger := log.With().
			Str("reqId", middleware.GetReqID(r.Context())).
			Logger()
		token := chi.URLParam(r, "token")
		sublogger.Info().Msgf("Token: %s", token)

		var req Req
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if token == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			var mutex = &sync.RWMutex{}
			mutex.RLock()
			if val, ok := s.Games[token]; ok {
				_, err := val.PostGame(req.Player, req.Row, req.Column)
				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				mutex.RUnlock()
				mutex.Lock()
				s.Games[token] = val
				mutex.Unlock()

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(val.GetGame())
				return
			}
			mutex.RUnlock()
		}
		w.WriteHeader(http.StatusNotFound)
	})
}

func (s *ServerAPI) DeleteGame(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sublogger := log.With().
			Str("reqId", middleware.GetReqID(r.Context())).
			Logger()
		token := chi.URLParam(r, "token")
		sublogger.Info().Msgf("Token: %s", token)

		var mutex = &sync.RWMutex{}
		mutex.RLock()
		if _, ok := s.Games[token]; ok {
			mutex.RUnlock()

			mutex.Lock()
			delete(s.Games, token)
			mutex.Unlock()
			deleted := map[string]string{"deleted": token}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deleted)
			return
		}
		mutex.RUnlock()
		w.WriteHeader(http.StatusNotFound)
	})
}

func (s *ServerAPI) HealthCheck(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"timestamp": "%s"}`, time.Now())))
	})
}

func (s *ServerAPI) Home(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"info": "%s"}`, "Welcome to the TicTacToe API. Please use the /game endpoint to start a new game.")))
	})
}
