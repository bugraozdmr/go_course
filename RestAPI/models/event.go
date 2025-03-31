package models

import (
	"rest/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(Name, Description, Location, DateTime, user_id) 
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	// TODO bunu fonksiyonlaştır sürekli if açmak istemiyorum artık check(err, message) gibi
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

func GetAllEvents() ([]Event, error) {
	// direkt parametresiz sesliceksen query yoksa exec hazirlancaksa felan
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET Name = ?, Description = ?, Location = ?, DateTime = ?
	WHERE id = ? 
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

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

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations (event_id,user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(e.ID, userId)

	return err
}