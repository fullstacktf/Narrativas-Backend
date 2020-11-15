package models

import (
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// CharacterSectionField : Structure
type CharacterSectionField struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	SectionID   uint   `gorm:"primary_key" json:"sectionId"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Value       string `gorm:"type:varchar(255);NOT NULL" json:"value"`
	Description string `json:"description"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

// TableName : Database table name map
func (CharacterSectionField) TableName() string {
	return "character_section_field"
}

// Insert : Inserts values into database
func (field *CharacterSectionField) Insert() {
	common.DB.Create(field)
}
