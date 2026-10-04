package models

import "example.com/event-mgt/db"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO
	users(name, email, password)
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}
