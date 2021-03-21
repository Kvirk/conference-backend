package db

import (
	"database/sql"
	"fmt"

	"github.com/kvirk/conference-backend/models"
)

var ErrDuplicate = fmt.Errorf("User with email exists")

func (db Database) AddUser(user *models.User) error {
	var id int
	var email string
	var firstName string
	var lastName string
	var password string
	var createdAt string
	doesUserExist, err := db.DoesUserByEmailExists(user.Email)

	if err != nil {
		return err
	}

	if !doesUserExist {
		return ErrDuplicate
	}

	query := `INSERT INTO users (email, first_name, last_name, password) VALUES ($1, $2, $3, $4) RETURNING id, email, first_name, last_name, password, created_at`
	err = db.Conn.QueryRow(query, user.Email, user.FirstName, user.LastName, user.Password).Scan(&id, &email, &firstName, &lastName, &password, &createdAt)
	if err != nil {
		return err
	}

	user.Id = id
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	user.Password = password
	user.CreatedAt = createdAt
	return nil
}

func (db Database) DoesUserByEmailExists(email string) (bool, error) {
	user := models.User{}

	query := `SELECT id, email, first_name, last_name, password, created_at FROM users WHERE email = $1;`
	row := db.Conn.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	return false, nil

}
