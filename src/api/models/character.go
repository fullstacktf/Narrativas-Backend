package models

import (
	"time"
)

type Character struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `gorm:"NOT NULL" json:"userId" binding:"required"`
	Name      string `gorm:"type:varchar(50)" json:"name"`
	Biography string `json:"biography"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Character_section struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
