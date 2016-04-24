package models

import (
	"time"
)

type Arks struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	Name       string    `sql:"size:255" json:"name"`
	ScreenName string    `sql:"size:255" json:"screen_name"`
	Ship       int       `json:"ship"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
