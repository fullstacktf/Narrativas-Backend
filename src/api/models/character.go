package models

import (
	"errors"
	"log"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

type Characters []CharacterData

type Character struct {
	ID               uint               `gorm:"primaryKey; ->; <-:create" json:"id"`
	UserID           uint               `gorm:"foreignKey; column:user_id" json:"userid"`
	Name             string             `gorm:"type:varchar(50)" json:"name" binding:"required"`
	Biography        string             `json:"biography" binding:"required"`
	Image            string             `gorm:"type:varchar(150)" json:"image" binding:"required"`
	UpdatedAt        time.Time          `json:"-"`
	CreatedAt        time.Time          `json:"-"`
	CharacterSection []CharacterSection `gorm:"-;foreignKey:CharacterID; references:ID" json:"sections,omitempty"`
}

type CharacterData struct {
	ID        uint   `json:"id"`
	Name      string `gorm:"type:varchar(50)" json:"name" binding:"required"`
	Biography string `json:"biography" binding:"required"`
	Image     string `gorm:"type:varchar(150)" json:"image" binding:"required"`
}

type Test struct {
	Character Character `json:"character"`
}

type CharacterSection struct {
	ID                    uint                    `gorm:"primaryKey;" json:"id"`
	CharacterID           uint                    `gorm:"column:character_id" json:"character_id"`
	Title                 string                  `gorm:"type:varchar(50)" json:"title"`
	UpdatedAt             time.Time               `json:"-"`
	CreatedAt             time.Time               `json:"-"`
	CharacterSectionField []CharacterSectionField `gorm:"-;foreignKey:SectionID;references:ID" json:"fields,omitempty"`
}

func (CharacterSection) TableName() string {
	return "character_section"
}

type CharacterSectionField struct {
	ID          uint      `gorm:"primaryKey; ->;<-:create" json:"id"`
	SectionID   uint      `gorm:"foreignKey" json:"section_id"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	Value       string    `gorm:"type:varchar(255);NOT NULL" json:"value"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (CharacterSectionField) TableName() string {
	return "character_section_field"
}

func (Character) TableName() string {
	return "actor"
}

func (test Test) Insert() error {
	// id, err := common.IsSignedIn(token)

	// if err != nil {
	// 	return false, err
	// }

	test.Character.UserID = 2
	test.Character.CreatedAt = time.Now()
	test.Character.UpdatedAt = time.Now()

	if len(test.Character.Name) > 50 || len(test.Character.Name) == 0 {
		return errors.New("character name too long")
	}

	common.DB.Debug().Create(&test.Character)
	for _, section := range test.Character.CharacterSection {
		common.DB.Debug().Create(&section)
		for _, field := range section.CharacterSectionField {
			common.DB.Debug().Create(&field)
		}
	}

	// common.DB.Debug().Create(&test.Character.CharacterSection)
	// common.DB.Debug().Create(&test.Character.CharacterSection.CharacterSectionField)

	// if result := common.DB.Debug().Create(&test.Character); result.Error != nil {
	// 	return errors.New("invalid data provided")
	// }
	return nil
}

func (character Character) Insert(token string) (bool, error) {
	id, err := common.IsSignedIn(token)

	if err != nil {
		return false, err
	}

	character.UserID = id
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

func (character Character) Delete(userid uint) error {
	common.DB.First(&character, character.ID)

	if character.Name == "" {
		return errors.New("Wrong character id")
	}

	if character.UserID != userid {
		return errors.New("Error deleting character")
	}

	common.DB.Delete(&character)
	return nil
}

func (characters *Characters) Get(id uint) error {
	rows, err := common.DB.Debug().
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

func (character Character) Get() (bool, error) {
	// TODO
	return true, nil
}
