package models

import (
	"errors"
	"log"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// Character : character BBDD structure
type Character struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `gorm:"foreign_key; column:user_id" json:"userid"`
	Name      string `gorm:"type:varchar(50)" json:"name" binding:"required"`
	Biography string `json:"biography" binding:"required"`
	Image     string `gorm:"type:varchar(150)" json:"image" binding:"required"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

// CharacterData : character API structure
type CharacterData struct {
	Name      string `gorm:"type:varchar(50)" json:"name" binding:"required"`
	Biography string `json:"biography" binding:"required"`
	Image     string `gorm:"type:varchar(150)" json:"image" binding:"required"`
}

// Characters : array of characters
type Characters []CharacterData

// TableName : Database table name map
func (Character) TableName() string {
	return "actor"
}

// Insert : Inserts character into bbdd
func (character Character) Insert() (bool, error) {
	// TODO: set ID depending the user
	character.UserID = 1
	character.CreatedAt = time.Now()
	character.UpdatedAt = time.Now()

	if len(character.Name) > 50 || len(character.Name) == 0 {
		return false, errors.New("character name too long")
	}

	if result := common.DB.Omit("Id").Create(&character); result.Error != nil {
		return false, errors.New("invalid data provided")
	}
	return true, nil
}

// Delete : Delete character from bbdd
func (character Character) Delete() (bool, error) {

	return true, nil
}

// Get : returns all characters of one user
func (characters *Characters) Get(id int) error {
	rows, err := common.DB.Debug().
		Model(&Character{}).
		Select(`actor.name,
				    actor.biography
					`).
		Joins("JOIN user on user.id = actor.user_id").
		Where("user.id = ?", id).
		Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		character := CharacterData{}
		err = common.DB.ScanRows(rows, &character)
		if err != nil {
			log.Println("error al bindear", err)
		} else {
			log.Printf("Character:%v\n", character)
		}
		*characters = append([]CharacterData(*characters), character)
	}
	return nil
}

// Get : Returns all information about one character
func (character Character) Get() (bool, error) {
	// TODO
	return true, nil
}
