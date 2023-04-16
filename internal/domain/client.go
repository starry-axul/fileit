package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Client model
type Client struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name     string         `gorm:"type:char(70);unique" json:"name"`
	ReadToken    string         `gorm:"type:char(36);unique" json:"-"`
	WriteToken     string         `gorm:"type:char(36);unique" json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (c *Client) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	if c.ReadToken == "" {
		c.ReadToken = uuid.New().String()
	}
	if c.WriteToken == "" {
		c.WriteToken = uuid.New().String()
	}
	return
}