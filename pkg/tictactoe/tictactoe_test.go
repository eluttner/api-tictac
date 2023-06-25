package tictactoe

import (
	"testing"

	tttoe "github.com/chrisfregly/tictactoe"
)

type Moves struct {
	Player string
	Row    int
	Column int
}

func TestWinX(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"O", 0, 1},
		{"X", 1, 1},
		{"O", 0, 2},
		{"X", 2, 2},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	}
	if *ttt.Game.GetWinner() != tttoe.X {
		t.Errorf("Winner should be X")
	}
}

func TestOverFlow(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"O", 0, 1},
		{"X", 1, 1},
		{"O", 0, 2},
		{"X", 2, 4},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: row and/or column is out of bounds"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}

func TestMoveXX(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"X", 0, 1},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: not X's turn"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}

func TestMoveNotEmpty(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"O", 0, 0},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: location 0,0 is not empty"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}

func TestMoveInvalidPlayer(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"Z", 0, 1},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: Invalid player Z, must be X or O"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}

func TestGameOver(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"O", 0, 1},
		{"X", 1, 1},
		{"O", 0, 2},
		{"X", 2, 2},
		{"O", 0, 1},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: game is already over"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}

func TestGameOverTie(t *testing.T) {
	moves := []Moves{
		{"X", 0, 0},
		{"O", 0, 1},
		{"X", 0, 2},

		{"O", 1, 1},
		{"X", 1, 0},
		{"O", 1, 2},

		{"X", 2, 1},
		{"O", 2, 0},
		{"X", 2, 2},

		{"O", 0, 0},
	}
	ttt := &TTT{}
	ttt.NewGame()

	for _, move := range moves {
		_, err := ttt.PostGame(move.Player, move.Row, move.Column)
		if err != nil {
			expected := "tictactoe: location 0,0 is not empty"
			if err.Error() != expected {
				t.Errorf("Expected: %s Got: %s", expected, err)
			}
		}
	}
}
