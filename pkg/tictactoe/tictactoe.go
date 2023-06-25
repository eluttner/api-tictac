package tictactoe

import (
	"fmt"

	ttt "github.com/chrisfregly/tictactoe"
	"github.com/google/uuid"
)

type TTTIf interface {
	NewGame() string
	GetGame(id string) TTTResponse
	PostGame(player string, row, column int) (error, TTTResponse)
}

type TTTResponse struct {
	Token  string     `json:"token"`
	Player string     `json:"next-player"`
	Winner string     `json:"winner"`
	Board  [][]string `json:"board"`
}

type TTTRequest struct {
	Token  string
	Player string
	Row    int
	Column int
}

type TTT struct {
	Token string
	Game  ttt.TicTacToe
}

// Create a new game and return the game token
func (t *TTT) NewGame() string {
	t.Token = uuid.New().String()
	t.Game = ttt.NewTicTacToe()
	return t.Token
}

func (t *TTT) PostGame(player string, row, column int) (TTTResponse, error) {
	p := ttt.X
	if player == "X" {
		p = ttt.X
	} else if player == "O" {
		p = ttt.O
	} else {
		return TTTResponse{}, fmt.Errorf("tictactoe: Invalid player %s, must be X or O", player)
	}
	err := t.Game.Move(p, row, column)
	if err != nil {
		return TTTResponse{}, err
	}
	return t.GetGame(), nil
}

// Get the current status of the game
func (t *TTT) GetGame() TTTResponse {
	tttResponse := TTTResponse{}

	tttResponse.Player = string(t.Game.GetTurn())

	if t.Game.GetWinner() != nil {
		tttResponse.Winner = tttResponse.Player
	} else {
		tttResponse.Winner = ""
	}

	tttResponse.Board = t.boardToString()

	tttResponse.Token = t.Token

	return tttResponse
}

func (t *TTT) boardToString() [][]string {
	board := t.Game.GetBoard()
	boardArray := make([][]string, len(board))
	for i := 0; i < len(board); i++ {
		boardArray[i] = make([]string, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == nil {
				boardArray[i][j] = "-"
			} else if *board[i][j] == ttt.X {
				boardArray[i][j] = string(ttt.X)
			} else {
				boardArray[i][j] = string(ttt.O)
			}
		}
	}
	return boardArray

	// board := t.Game.GetBoard()
	// boardString := "\n"
	// for i := 0; i < len(board); i++ {
	// 	for j := 0; j < len(board[i]); j++ {
	// 		if board[i][j] == nil {
	// 			boardString += "-"
	// 		} else if *board[i][j] == ttt.X {
	// 			boardString += string(ttt.X)
	// 		} else {
	// 			boardString += string(ttt.O)
	// 		}
	// 	}
	// 	boardString += "\n"
	// }
	// return boardString
}
