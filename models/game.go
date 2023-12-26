package models

import "github.com/google/uuid"

type Game struct {
	Id    uuid.UUID
	Board Board
}
