package models

import (
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
	ID            uint            `gorm:"primaryKey; ->; <-:create" json:"id,omitempty"`
	StoryID       int             `gorm:"column:story_id; foreignKey:story_id" json:"story_id`
	Title         string          `gorm:"column:title" json:"title"`
	Description   string          `gorm:"column:description" json:"description"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       `json:"-"`
	EventRelation []EventRelation `gorm:"column: event_relation; foreignKey:initial_event; references:ID" json:"events_relations,omitempty"`
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

// TableName : Database table name map
func (Story) TableName() string {
	return "story"
}

// TableName : Database table name map
func (Event) TableName() string {
	return "event"
}

// TableName : Database table name map
func (EventRelation) TableName() string {
	return "event_relation"
}

// Stories : New type. Arrays of stories
type Stories []Story

// Get : Get all the stories in the DB

// ******************************* FUNCIONA CORRECTAMENTE ***************************************//
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
//  ******************************************************* FUNCIONA CORRECTAMENTE ************************************************//
func (s *Story) Get(id string) error {

	storyID, _ := strconv.Atoi(id)

	common.DB.
		Model(&Story{}).
		Preload("Event.EventRelation").
		Select(`story.id, 
				story.title, 
				story.image,
				event.title,
				event.description,
				event_relation.initial_event,
				event_relation.final_event`).
		Joins("JOIN user ON user.id = story.user_id").
		Joins("INNER JOIN event").
		Joins("INNER JOIN event_relation").
		Where("user.id = ?  AND story.id = ? AND event.story_id = ?", *&s.UserID, storyID, storyID).
		Find(&s)

	return nil
}

//  ******************************************************* ERROR : NO FUNCIONA ************************************************//

// Insert : Add a new story
func (s *Story) Insert() error {
	result := common.DB.Debug().Create(s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//  ******************************************************* ERROR : NO FUNCIONA ************************************************//

// InsertEvent : Add a new event
func (e *Event) InsertEvent() error {
	result := common.DB.Debug().Create(e)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//  ******************************************************* ERROR : NO FUNCIONA ************************************************//

// Update : Update a story in the database
func (s *Story) Update() error {
	result := common.DB.Debug().Save(s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//  ******************************************************* ERROR : NO FUNCIONA ************************************************//

// Delete : Delete a story in the database
// El delete funcionaba solo con la estrutura story pero ahora se queja y dice que necesita un where
// Error : WHERE conditions required
func (s *Story) Delete(id string) error {
	common.DB.Debug().First(&s, id)
	result := common.DB.Debug().Delete(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
