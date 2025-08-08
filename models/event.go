package models

import (
	"time"
)

type Event struct {
	ID uint `gorm:"primarykey" json:"id"`

	Name        string    `json:"name" binding:"required" gorm:"not null"`
	Description string    `json:"description" binding:"required" gorm:"not null"`
	Location    string    `json:"location" binding:"required" gorm:"not null"`
	DateTime    time.Time `json:"datetime" binding:"required" gorm:"not null"`
	UserID      int       `json:"user_id" gorm:"not null;default:1"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}