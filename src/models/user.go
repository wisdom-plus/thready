package models

import (
	"thready/src/db"
	"thready/src/utils"
	"time"
)

type User struct {
	ID 							int       `db:"id"`
	Name 						string    `db:"name"`
	PasswordHash 		string    `db:"password_hash"`
	CreatedAt 			time.Time `db:"created_at"`
}

func CreateUser(name, password string) (int, error) {
	passwordHash, passwordErr := utils.HashPassword(password)
	if passwordErr != nil {
		return 0, passwordErr
	}

	var id int
	err := db.DB.QueryRow("INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id", name, passwordHash).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func FindUserByLogin(name, password string) (*User, error) {
	user := &User{}
	err := db.DB.QueryRow("SELECT id, name, created_at FROM users WHERE name = $1 AND password_hash = $2", name, password).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetCurrentUser(id int) (*User, error) {
	user := &User{}
	err := db.DB.QueryRow("SELECT id, name, created_at FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
