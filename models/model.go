package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model definition
type Model struct {
	Id        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
