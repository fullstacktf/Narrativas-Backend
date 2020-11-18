package models

import "time"

// CharacterSection : Character section structure
type CharacterSection struct {
	ID                    uint `gorm:"primaryKey" json:"id"`
	CharacterSectionField []CharacterSectionField
	UpdatedAt             time.Time
	CreatedAt             time.Time
}
