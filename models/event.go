package models

import "time"

type Event struct { 
	ID          int
	Name        string    `json:"name" binding:"required"`
		Description string    `json:"description" binding:"required"`
		Location    string    `json:"location" binding:"required"`
		DateTime    time.Time `json:"datetime" binding:"required"`
		UserID      int
}


var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event { 
	return events
} 