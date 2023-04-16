package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Client model
type File struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	FileName     string         `gorm:"type:char(100);unique" json:"filename"`
	URLReference		string `gorm:"type:char(300)" json:"-"`
	Type    string         `gorm:"type:char(20)" json:"type"`
	Client 		 Client     `json:"client"`
	ClientID     string         `gorm:"type:char(36)" json:"client_id"`
	Private bool				`json:"private"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (c *File) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return
}