package models

type Player uint8

const (
	NOT_SET Player = iota
	X
	O
)

func (p Player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return ""
	}
}

type Board struct {
	CurrentPlayer Player
	Cells         []Player
}
