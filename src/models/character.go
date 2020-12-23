package models

import (
	"errors"
	"time"

	"github.com/fullstacktf/Narrativas-Backend/common"
)

type CharacterData struct {
	Character Character `json:"character"`
}

type Characters []Character

type Character struct {
	ID               uint               `gorm:"primaryKey; ->; <-:create" json:"id"`
	UserID           uint               `gorm:"foreignKey; column:user_id" json:"-"`
	Name             string             `gorm:"type:varchar(50)" json:"name" binding:"required"`
	Biography        string             `json:"biography" binding:"required"`
	Image            string             `gorm:"type:varchar(150)" json:"image" binding:"required"`
	UpdatedAt        time.Time          `json:"-"`
	CreatedAt        time.Time          `json:"-"`
	CharacterSection []CharacterSection `gorm:"foreignKey:CharacterID; references:ID" json:"sections,omitempty"`
}

type CharacterSection struct {
	ID                    uint                    `gorm:"primaryKey;" json:"id,omitempty"`
	CharacterID           uint                    `gorm:"column:character_id" json:"character_id"`
	Title                 string                  `gorm:"type:varchar(50)" json:"title"`
	UpdatedAt             time.Time               `json:"-"`
	CreatedAt             time.Time               `json:"-"`
	CharacterSectionField []CharacterSectionField `gorm:"foreignKey:SectionID;references:ID" json:"fields,omitempty"`
}

type CharacterSectionField struct {
	ID          uint      `gorm:"primaryKey; ->;<-:create" json:"id,omitempty"`
	SectionID   uint      `gorm:"foreignKey" json:"section_id"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	Value       string    `gorm:"type:varchar(255);NOT NULL" json:"value"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (Character) TableName() string {
	return "actor"
}

func (CharacterSection) TableName() string {
	return "character_section"
}

func (CharacterSectionField) TableName() string {
	return "character_section_field"
}

func (character *Character) Insert() error {

	if result := common.DB.Create(&character); result.Error != nil {
		return errors.New("invalid data provided")
	}
	return nil
}

func (character *Character) Delete(userid uint) error {
	common.DB.First(&character, character.ID)

	if character.Name == "" || character.UserID != userid {
		return errors.New("error deleting character")
	}

	common.DB.Delete(&character)
	return nil
}

func (characters *Characters) Get(id uint) error {
	rows, err := common.DB.
		Model(&Character{}).
		Select(`actor.id,
						actor.name,
						actor.biography,
						actor.image
					`).
		Joins("JOIN user on user.id = actor.user_id").
		Where("user.id = ?", id).
		Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		character := Character{}
		err = common.DB.ScanRows(rows, &character)
		if err != nil {
			return err
		}
		*characters = append([]Character(*characters), character)
	}
	return nil
}

func (character *Character) Get(userid uint) error {

	common.DB.
		Model(&Character{}).
		Preload("CharacterSection.CharacterSectionField").
		Select(`actor.name,
					actor.user_id,
					actor.biography,
					actor.image,
					character_section.title,
					character_section_field.name,
					character_section_field.description,
					character_section_field.value
					`).
		Joins("INNER JOIN character_section").
		Joins("INNER JOIN character_section_field").
		Where("actor.id = ? AND actor.id = character_section.character_id AND character_section.id = character_section_field.section_id", character.ID).
		Find(&character)

	if character.Name == "" || character.UserID != userid {
		common.DB.
			Model(&Character{}).
			Select(`actor.name,
					actor.user_id,
					actor.biography,
					actor.image
					`).
			Find(&character)

		if character.UserID != userid {
			return errors.New("character not found")
		}
	}

	return nil
}

func (character Character) Update() error {
	var deleteCharacter Character
	deleteCharacter.ID = character.ID

	if err := deleteCharacter.Delete(character.UserID); err != nil {
		return err
	}

	if result := common.DB.Create(&character); result.Error != nil {
		return errors.New("invalid data provided")
	}
	return nil
}

func (section *CharacterSection) Insert() error {

	if result := common.DB.Create(&section); result.Error != nil {
		return errors.New("invalid data provided")
	}
	return nil
}

func (field *CharacterSectionField) Insert() error {

	if result := common.DB.Create(&field); result.Error != nil {
		return errors.New("invalid data provided")
	}
	return nil
}
