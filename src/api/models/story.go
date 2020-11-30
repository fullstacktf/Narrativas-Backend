package models

import (
	"log"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// Story : Structure
type Story struct {
	//gorm.Model
	ID             int       `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	UserID         uint      `gorm:"column:user_id; foreignKey:user_id" json:"user_id"`
	InitialEventID uint      `gorm:"column:initial_event_id;foreignKey:initial_event_id;default:null" json:"initial_event_id,omitempty"`
	Image          string    `gorm:"type:varchar(150);column:image;NOT NULL" json:"image"`
	Title          string    `gorm:"type:varchar(150);column:title" json:"title" binding:"required"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	//Description string    `gorm:"type:string; NOT NULL; column:description" json:"description" binding:"required"`
	//User        User      `gorm:"foreignKey:user_id; references:user_id"`
}

// TableName : Database table name map
func (Story) TableName() string {
	return "story"
}

// Stories : New type. Arrays of stories
type Stories []Story

// Get : Get all the stories in the DB
func (s *Stories) Get() error {
	rows, err := common.DB.
		Model(&Story{}).
		Select(`story.id, story.title`).
		Joins("JOIN user ON user.id = story.user_id").
		Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		story := Story{}
		err = common.DB.ScanRows(rows, &story)
		if err != nil {
			log.Println("error al bindear", err)
		} else {
			log.Printf("Story:%v\n", story)
		}
		*s = append([]Story(*s), story)
	}
	return nil
}

// Get : Get only one story through the id
func (s *Story) Get(id string) error {
	common.DB.
		Model(&Story{}).
		Select(`story.id, story.title`).
		Joins("JOIN user ON user.id = story.user_id").
		Where("story.id = ?", &id).
		Find(&s)

	return nil
}

// Insert : Add a new story
func (s *Story) Insert() error {
	result := common.DB.Debug().Create(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update : Update a story in the database
func (s *Story) Update() error {
	result := common.DB.Debug().Save(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete : Delete a story in the database
func (s *Story) Delete() error {
	result := common.DB.Debug().Delete(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
