package models

import (
	"log"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// Story : Structure
type Story struct {
	// gorm.Model
	ID        uint   `gorm:"column:id;primaryKey" json:"id"`
	IDUser    uint   `gorm:"column:user_id;type:uint; NOT NULL;"`
	Title     string `gorm:"column:title;type:varchar(255); NOT NULL" json:"title" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeleteAt  time.Time
	// description string `gorm:"type:string; NOT NULL" json:"title" binding:"required"`
	User `gorm:"foreignKey:user_id;references:user_id"`
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
					user.Username`).
		Joins("JOIN user on user.id = story.user_id").
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
		Select([]string{"title", "user.Username"}).
		Where("story.idUser = ?", &id).
		Joins("JOIN user on user.id = story.idUser").
		Row()

	story := &Story{}

	row.Scan(&story.Title, &story.User.Username)

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
