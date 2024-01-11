package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func NewGame() *Game {
	board := make([]int32, 9)
	for i := range board {
		board[i] = 0
	}
	game := new(Game)
	game.Id = uuid.New()
	game.CurrentPlayer = 1
	game.Cells = board
	return game
}

type Game struct {
	Model
	CurrentPlayer int32         `gorm:"type:integer; default:0"`
	Cells         pq.Int32Array `gorm:"type:integer[]"`
	Winner        int32         `gorm:"type:integer"`
	Winners       pq.Int32Array `gorm:"type:integer[]"`
}

func (g *Game) SwitchPlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
	}
}

func (g *Game) UpdateCell(index int) {
	g.Cells[index] = int32(g.CurrentPlayer)
}

func (g *Game) IsCellWinner(index int) bool {
	for _, winner := range g.Winners {
		if i := int32(index); winner == i {
			return true
		}
	}
	return false
}

func (g *Game) DisplayCellValue(index int) string {
	switch g.Cells[index] {
	case 1:
		return "X"
	case 2:
		return "O"
	default:
		return ""
	}
}

func (g *Game) DisplayWinnerValue() string {
	switch g.Winner {
	case 1:
		return "Player X Wins"
	case 2:
		return "Player O Wins"
	case 3:
		return "It's a Draw"
	default:
		return ""
	}
}

var winners = [][]int32{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func (g *Game) FindWinner() {
	p := g.CurrentPlayer
	for _, v := range winners {
		tic, tac, toe := v[0], v[1], v[2]
		if g.Cells[tic] == p && g.Cells[tac] == p && g.Cells[toe] == p {
			g.Winner = p
			g.Winners = v
			return
		}
	}
	for _, v := range g.Cells {
		if v == 0 {
			return
		}
	}
	g.Winner = 3 // Draw
}
