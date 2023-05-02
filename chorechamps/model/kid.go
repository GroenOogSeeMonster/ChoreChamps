package model

import (
	"time"
)

type Kid struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	WhatsApp  string    `gorm:"type:varchar(255);not null" json:"whatsApp"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}