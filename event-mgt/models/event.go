package models

import (
	"time"

	"example.com/event-mgt/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func Save(e *Event) error {
	query := `INSERT INTO 
	events(title, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime.Format(time.RFC3339), e.UserID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var datetimeStr string
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &datetimeStr, &event.UserID)
		if err != nil {
			return nil, err
		}
		event.DateTime, err = time.Parse(time.RFC3339, datetimeStr)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	var datetimeStr string
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &datetimeStr, &event.UserID)
	if err != nil {
		return nil, err
	}
	event.DateTime, _ = time.Parse(time.RFC3339, datetimeStr)
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events 
	SET title = ?, description = ?, location = ?, datetime = ? 
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.DateTime.Format(time.RFC3339), event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (event Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)
	return err
}
