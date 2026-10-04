package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func Save(e Event) {
	// TODO: Save event to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
