package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//User model
type User struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserName     string         `gorm:"type:char(70);unique" json:"username"`
	FirstName    string         `gorm:"type:char(30)" json:"firstname"`
	LastName     string         `gorm:"type:char(30)" json:"lastname"`
	Password     string         `json:"password,omitempty"`
	Email        string         `gorm:"type:char(70)" json:"email"`
	Token        string         `json:"token" gorm:"type:char(36);not null"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}