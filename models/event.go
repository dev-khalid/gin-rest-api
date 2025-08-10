package models

import (
	"time"
)

type EventAccessType string

const (
	Free    EventAccessType = "free"
	Premium EventAccessType = "premium"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "credit_card"
	Cash       PaymentMethod = "cash"
)

type EventPricing struct {
	ID              uint            `gorm:"primarykey" json:"id"`
	Price           float64         `gorm:"not null;default:0" json:"price"`
	Currency        string          `gorm:"not null;default:'USD'" json:"currency"`
	EventAccessType EventAccessType `gorm:"type:enum('free','premium');not null;default:'free'" json:"event_access_type"`
	PaymentMethod   PaymentMethod   `gorm:"type:enum('credit_card','cash');not null;default:'credit_card'" json:"payment_method"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Event struct {
	ID uint `gorm:"primarykey" json:"id"`

	Name           string    `json:"name" binding:"required" gorm:"not null"`
	Description    string    `json:"description" binding:"required" gorm:"not null"`
	Location       string    `json:"location" binding:"required" gorm:"not null"`
	DateTime       time.Time `json:"datetime" binding:"required" gorm:"not null"`
	UserID         int       `json:"user_id" gorm:"not null;default:1"`
	EventPricingId uint
	EventPricing   EventPricing `gorm:"constraint:OnDelete:CASCADE;" json:"event_pricing"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
