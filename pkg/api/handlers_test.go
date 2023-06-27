package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eluttner/api-tictac/pkg/tictactoe"
	"github.com/go-chi/chi/v5"
)

type Moves struct {
	Move     string
	Expected string
}

func TestHealthCheck(t *testing.T) {
	t.Run("TestHealthCheck", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/health-check", nil)
		if err != nil {
			t.Fatal(err)
		}
		ctx := context.Background()
		s := &ServerAPI{}
		s.ConfigureServer(ctx)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.HealthCheck(ctx))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"timestamp":`
		if !strings.HasPrefix(rr.Body.String(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestHome(t *testing.T) {
	t.Run("TestHome", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		ctx := context.Background()
		s := &ServerAPI{}
		s.ConfigureServer(ctx)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.Home(ctx))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"info":`
		if !strings.HasPrefix(rr.Body.String(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestGame(t *testing.T) {
	//t.Run("TestGameWin", func(t *testing.T) {
	ctx := context.Background()
	s := &ServerAPI{}
	s.ConfigureServer(ctx)

	expected_initial := `{"token": "","next-player": "X","winner": "","board": [["-","-","-"],["-","-","-"],"-","-","-"]]}`
	moves := []Moves{
		{`{"player": "X","row": 0,"column": 0}`, `{"token": "TOKEN","next-player": "O","winner": "","board": [["X","-","-"],["-","-","-"],"-","-","-"]]}`},
		{`{"player": "O","row": 0,"column": 1}`, `{"token": "TOKEN","next-player": "X","winner": "","board": [["-","O","-"],["-","-","-"],"-","-","-"]]}`},
		{`{"player": "X","row": 1,"column": 1}`, `{"token": "TOKEN","next-player": "O","winner": "","board": [["-","X","-"],["-","-","-"],"-","-","-"]]}`},
		{`{"player": "O","row": 0,"column": 2}`, `{"token": "TOKEN","next-player": "X","winner": "","board": [["-","-","O"],["-","-","-"],"-","-","-"]]}`},
		{`{"player": "X","row": 2,"column": 2}`, `{"token": "TOKEN","next-player": "O","winner": "X","board": [["-","-","X"],["-","-","-"],"-","-","-"]]}`},
	}

	req, err := http.NewRequest("GET", "/game", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.GetGame(ctx))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	check, err := verifyResponse(rr.Body.String(), expected_initial, false)
	if err != nil {
		t.Fatal(err)
	}
	if !check {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected_initial)
	}
	token, err := getToken(rr.Body.String())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)
	for _, move := range moves {
		body := []byte(move.Move)
		r, err := http.NewRequest("POST", fmt.Sprintf("/game/%s/move", token), bytes.NewBuffer(body))

		chiCtx := chi.NewRouteContext()
		req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		// Add the key/value to the context.
		chiCtx.URLParams.Add("token", fmt.Sprintf("%v", token))

		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(s.PostGame(ctx))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := strings.Replace(move.Expected, "TOKEN", token, 1)
		check, err = verifyResponse(rr.Body.String(), expected, true)
		if err != nil {
			t.Fatal(err)
		}
		if !check {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
		token, err = getToken(rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}
	}
	//})
}

func getToken(resp string) (string, error) {
	var respObj tictactoe.TTTResponse
	err := json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		return "", err
	}
	return respObj.Token, nil

}
func verifyResponse(resp, expected string, ignoreToken bool) (bool, error) {
	var respObj tictactoe.TTTResponse
	var expObj tictactoe.TTTResponse
	err := json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(resp), &expObj)
	if err != nil {
		return false, err
	}

	if respObj.Token != expObj.Token {
		return false, nil
	}

	if respObj.Player != expObj.Player {
		return false, nil
	}

	if respObj.Winner != expObj.Winner {
		return false, nil
	}

	for i := 0; i < len(respObj.Board); i++ {
		for j := 0; j < len(respObj.Board[i]); j++ {
			if respObj.Board[i][j] != expObj.Board[i][j] {
				return false, nil
			}
		}
	}

	return true, nil
}
