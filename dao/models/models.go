package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        string         `gorm: PRIMARY_KEY; json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Todo struct {
	Base            Model
	ExpiryTime      time.Time `gorm:"not null"`
	Title           string    ``
	Description     string    ``
	PercentComplete int32     `gorm:"check:0-100"`
}
