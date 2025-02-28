package models

import "restapi.com/m/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	// "?", protege al codigo contra inyecciones sql
	query := `
	INSERT INTO users(email, password) 
	VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//cuando todo el codigo se halla ejecutado cierro la conexion
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}
