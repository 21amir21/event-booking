package models

import "time"

type Event struct {
	ID          int
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e Event) Save() {
	// TODO: we will store it to the database
	events = append(events, e)
}

func GetAllEevents() []Event {
	return events
}
