package models

import (
	"time"

	"example.com/eventer-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, user_id)
		VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return err
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, created_at = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.CreatedAt, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userID int64) error {
	query := `INSERT INTO registrations(user_id, event_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, e.ID)
	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := `DELETE FROM registrations WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, created_at, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.CreatedAt, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, created_at, user_id FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.CreatedAt, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}
