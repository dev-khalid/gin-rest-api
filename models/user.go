package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primarykey"`

	Name     string `json:"name" binding:"required" gorm:"not null"`
	Email    string `json:"email" binding:"required" gorm:"not null;unique"`
	Password string `json:"password" binding:"required" gorm:"not null"`
	Role     Role   `json:"role" gorm:"not null;default:'User';type:enum('Admin', 'User')"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
