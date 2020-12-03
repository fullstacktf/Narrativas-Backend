package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

// Story : Structure
type Story struct {
	//gorm.Model
	ID             int       `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	UserID         uint      `gorm:"column:user_id; foreignKey:user_id" json:"user_id"`
	InitialEventID uint      `gorm:"column:initial_event_id;foreignKey:initial_event_id;default:null" json:"initial_event_id,omitempty"`
	Image          string    `gorm:"column:image;NOT NULL" json:"image"`
	Title          string    `gorm:"column:title" json:"title" binding:"required"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	//Description string    `gorm:"type:string; NOT NULL; column:description" json:"description" binding:"required"`
	Event []Event `gorm:"foreignKey:StoryID; references:ID" json:"events,omitempty"`
}

//Event : Structure

type Event struct {
	ID          uint      `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	StoryID     uint      `gorm:"column:story_id; foreignKey:story_id" json:"story_id`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	// EventRelationInitial []EventRelation `gorm:"column: father_event; foreignKey:initial_event; references:ID" json:"events,omitempty"`
	// EventRelationFinal   []EventRelation `gorm:"column: children_event; foreignKey:final_event; references:ID" json:"events,omitempty"`
}

// EventRelation : Structure

type EventRelation struct {
	ID           uint `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	InitialEvent uint `gorm:"column:initial_event; foreignKey:initial_event" json:"initial_event`
	FinalEvent   uint `gorm:"column:final_event; foreignKey:final_event;" json:"final_event"`
	CreatedAt    uint `json:"-"`
	UpdatedAt    uint `json:"-"`
}

// Story : Database table name map
func (Story) TableName() string {
	return "story"
}

// Event : Database table name map
func (Event) TableName() string {
	return "event"
}

// EventRelation : Database table name map
func (EventRelation) TableName() string {
	return "event_relation"
}

// Stories : New type. Arrays of stories
type Stories []Story

// Get : Get all the stories in the DB
func (s *Stories) Get(userID uint) error {
	rows, err := common.DB.
		Model(&Story{}).
		Select(`story.id,
		story.image, 
		story.title`).
		Joins("JOIN user ON user.id = story.user_id").
		Where("user.id = ?", userID).
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

	storyID, _ := strconv.Atoi(id)

	fmt.Println("ID: \n", storyID)
	fmt.Println("user_id: \n", *(&s.UserID))
	common.DB.
		Model(&Story{}).
		Preload("Event").
		Select(`story.id, 
				story.title, 
				story.image,
				event.title,
				event.description`).
		Joins("JOIN user ON user.id = story.user_id").
		Joins("INNER JOIN event").
		Where("user.id = ?  AND story.id = ? AND event.story_id = ?", *(&s.UserID), storyID, storyID). //Corregir
		Find(&s)

	return nil
}

// Insert : Add a new story
func (s *Story) Insert(userID uint) error {
	if userID == *(&s.UserID) {
		result := common.DB.Debug().Create(s)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// Update : Update a story in the database
func (s *Story) Update(userID uint) error {
	if userID == *(&s.UserID) {
		result := common.DB.Debug().Save(s)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// Delete : Delete a story in the database
func (s *Story) Delete(userID uint) error {
	if userID == *(&s.UserID) {
		result := common.DB.Debug().Delete(s)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
