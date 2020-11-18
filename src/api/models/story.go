package models

import (
	"log"
	"time"

	"gorm.io/gorm"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// Story : Structure
type Story struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	idUser    uint   `gorm:"type:uint; NOT NULL" json:"user_id" binding:"required"`
	title     string `gorm:"type:varchar(255); NOT NULL" json:"title" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      `gorm:"foreignKey:IdUser;references:idUser"`
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
		Select(`story.title,
					user.name`).
		Joins("JOIN user on user.id = story.idUser").
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
	row := common.DB.
		Model(&Story{}).
		Select([]string{"title", "user.name"}).
		Where("story.idUser = ?", &id).
		Joins("JOIN user on user.id = story.idUser").
		Row()

	story := &Story{}

	row.Scan(&story.title, &story.User.Username)

	return nil
}

// Insert : Add a new story
func (s *Story) Insert() error {
	result := common.DB.Debug().Save(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete : Delete a story in the database

// func (s *Story) Delete() error {
// 	var
// }
