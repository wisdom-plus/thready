package models

import (
	"thready/src/db"
	"time"
)

type Thread struct {
    ID        int       `db:"id"`
    Title     string    `db:"title"`
    UserID    int       `db:"user_id"`
    CreatedAt time.Time `db:"created_at"`
}

func GetAllThreads() ([]Thread, error) {
    threads := []Thread{}
    err := db.DB.Select(&threads, "SELECT id, title, created_at FROM threads ORDER BY id DESC")
    return threads, err
}

func CreateThread(title string, userID int) (int, error) {
    var id int
    err := db.DB.QueryRow("INSERT INTO threads (title, user_id) VALUES ($1, $2) RETURNING id", title, userID).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, err
}

func FindThreadByID(id int) (*Thread, error) {
    thread := Thread{}
    err := db.DB.Get(&thread, "SELECT id, title, created_at FROM threads WHERE id = $1", id)
    if err != nil {
        return nil, err
    }
    return &thread, nil
}
