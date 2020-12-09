package models

import (
	"errors"
	"log"
	"time"

	common "github.com/fullstacktf/Narrativas-Backend/common"
)

type Story struct {
	ID             uint      `gorm:"primaryKey;" json:"id,omitempty"`
	UserID         uint      `gorm:"column:user_id; foreignKey:user_id" json:"-"`
	InitialEventID uint      `gorm:"column:initial_event_id;foreignKey:initial_event_id;default:null" json:"initial_event_id,omitempty"`
	Image          string    `gorm:"column:image;NOT NULL" json:"image,omitempty" binding:"required"`
	Title          string    `gorm:"column:title" json:"title,omitempty" binding:"required"`
	Description    string    `gorm:"type:string; NOT NULL; column:description" json:"description,omitempty" binding:"required"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Event          []Event   `gorm:"foreignKey:StoryID; references:ID" json:"events,omitempty"`
}

type Event struct {
	ID            uint             `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	StoryID       uint             `gorm:"column:story_id;foreignKey:story_id;references:ID" json:"story_id"`
	Title         string           `gorm:"column:title" json:"event_title"`
	Description   string           `gorm:"column:description" json:"event_description"`
	CreatedAt     time.Time        `json:"-"`
	UpdatedAt     time.Time        `json:"-"`
	EventRelation []*EventRelation `gorm:"column:event_relation; foreignKey:initial_event; references:ID" json:"event_relation,omitempty"`
}

type EventRelation struct {
	ID           uint      `gorm:"primaryKey; <-:create" json:"id,omitempty"`
	InitialEvent uint      `gorm:"column:initial_event; foreignKey:initial_event" json:"initial_event"`
	FinalEvent   uint      `gorm:"column:final_event; foreignKey:final_event;" json:"final_event"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func (Story) TableName() string {
	return "story"
}

func (Event) TableName() string {
	return "event"
}

func (EventRelation) TableName() string {
	return "event_relation"
}

type Stories []Story

func (s *Stories) Get(userID uint) error {
	rows, err := common.DB.
		Model(&Story{}).
		Select(`story.id,
		story.image, 
		story.title,
		story.description`).
		Joins("JOIN user ON user.id = story.user_id").
		Where("user.id = ?", userID).
		Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		story := Story{}

		if err := common.DB.ScanRows(rows, &story); err == nil {
			log.Println("error al bindear", err)
		} else {
			log.Printf("Story:%v\n", story)
		}
		*s = append([]Story(*s), story)
	}
	return nil
}

func (story *Story) Get(storyID uint) error {

	common.DB.
		Model(&Story{}).
		Preload("Event.EventRelation").
		Select(`
				event_relation.initial_event,
				event_relation.final_event,	
				event.title,
				event.description,
				story.id,
				story.image,
				story.description,
				story.title
		`).
		Joins("JOIN user ON user.id = story.user_id").
		Joins("INNER JOIN event").
		Joins("INNER JOIN event_relation").
		Where("user.id = ? AND event.story_id = ? AND story.id = ? AND event.id = event_relation.initial_event", story.UserID, storyID, storyID).
		Find(&story)

	if story.Title == "" {
		userid := story.UserID

		common.DB.
			Model(&Story{}).
			Preload("Event").
			Select(`story.id,
				story.image,
				story.description,
				story.title`).
			Joins("JOIN user ON user.id = story.user_id").
			Where("user.id = ? AND story.id = ?", userid, storyID).
			Find(&story)
	}

	return nil
}

func (story *Story) Insert() error {
	result := common.DB.Create(story)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (event *Event) Insert() error {
	result := common.DB.Omit("event_relation").Create(event)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (relation *EventRelation) Insert() error {
	result := common.DB.Create(relation)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (story *Story) Delete(userid uint) error {
	common.DB.First(&story, story.ID)

	if story.Title == "" || story.UserID != userid {
		return errors.New("error deleting story")
	}

	common.DB.Delete(&story)
	return nil
}
