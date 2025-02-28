package models

import (
	"errors"

	"restapi.com/m/db"
	"restapi.com/m/utils"
)

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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (u User) ValidateCredentials() error {

	// "?", protege al codigo contra inyecciones sql
	query := `
	SELECT password
	FROM users
	WHERE email = ? 
	`

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&retrivedPassword)

	if err != nil {
		return errors.New("credenciales invalidas")
	}

	passwordIsValid := utils.CheckPasswordHash(retrivedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("credenciales invalidas")
	}

	return nil
}
