package models

import (
	"time"
)

// Character : structure
type Character struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `gorm:"NOT NULL" json:"userId" binding:"required"`
	Name      string `gorm:"type:varchar(50)" json:"name"`
	Biography string `json:"biography"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

// TableName : Database table name map
func (Character) TableName() string {
	return "character"
}
