package models

import (
	"time"

	"restapi.com/m/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {

	// "?", protege al codigo contra inyecciones sql
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()
	return err
}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer rows.Close()

	var events []Event

	for rows.Next() {

		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {

	query := "SELECT * FROM events  WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {

	query := `
	UPDATE events 
	SET name =?, description=?,  location=?, dateTime=?
	WHERE id=?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err

}

func (e Event) Delete() error {

	query := `
	DELETE 
	FROM events 
	WHERE id=?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e *Event) Register(userId int64) error {

	// "?", protege al codigo contra inyecciones sql
	query := `
	INSERT INTO registrations(event_id, user_id) 
	VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {

	query := `
	DELETE 
	FROM registrations 
	WHERE event_id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
