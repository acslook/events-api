package model

import (
	"events-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
			  VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	query := `UPDATE events SET name=?, description=?, location=?, date_time=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Register(userID int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := `DELETE FROM registrations WHERE event_id=? AND user_id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	return err
}
