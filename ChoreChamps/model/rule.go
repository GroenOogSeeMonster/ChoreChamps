package model

import (
	"time"
)

type Rule struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Score       int       `gorm:"type:int;not null" json:"score"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}