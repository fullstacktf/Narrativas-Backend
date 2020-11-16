package models

import (
	"log"

	"gorm.io/gorm"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// User : Structure
type Story struct {
	gorm.Model
	ID     uint   `gorm:"primary_key" json:"id"`
	idUser uint   `gorm:"type:uint; NOT NULL" json:"user_id" binding:"required"`
	title  string `gorm:"type:varchar(255); NOT NULL" json:"title" binding:"required"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	User `gorm:"foreignKey:IdUser;references:idUser"`
}

// TableName : Database table name map
func (Story) TableName() string {
	return "story"
}

type Stories []Story

func (s *Stories) Get() error {
	rows, err := common.DB.
		Model(&Story{}).
		Select(`story.title,
					user.name`).
		Joins("JOIN user on user.Username = story.idUser").
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

func Get(id int) error {
	row := common.DB.
		Model(&Story{}).
		Select([]string{"title", "user.name"}).
		Where("story.idUser = ?", &id).
		Joins("JOIN user on user.name = story.idUser").
		Row()

	story := &Story{}

	row.Scan(&story.title, &story.User.Username)
}
