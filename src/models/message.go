package models

import (
	"thready/src/db"
	"time"
)

type Message struct {
	ID				int       `db:"id"`
	ThreadID		int       `db:"thread_id"`
	Content			string    `db:"content"`
	CreatedAt		time.Time `db:"created_at"`
}

func FindMessageByThreadID(threadID int) ([]Message, error) {
	messages := []Message{}
	err := db.DB.Select(&messages, "SELECT id, thread_id, content, created_at FROM messages WHERE thread_id = $1 ORDER BY id DESC", threadID)
	return messages, err
}

func CreateMessage(threadID int, content string) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO messages (thread_id, content) VALUES ($1, $2) RETURNING id", threadID, content).Scan(&id)
	return id, err
}

