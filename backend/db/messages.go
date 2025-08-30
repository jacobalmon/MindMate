package db

import (
	"context"
	"log"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Role      string    `json:"role"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertMessage(userID int, role, text string) error {
	_, err := Pool.Exec(context.Background(),
		"INSERT INTO messages (user_id, role, text) VALUES ($1, $2, $3)", userID, role, text)
	if err != nil {
		log.Println("InsertMessage error", err)
	}
	return err
}

func GetLastMessages(userID, limit int) ([]Message, error) {
	rows, err := Pool.Query(context.Background(),
		"SELECT id, role, text, created_at FROM messages WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2",
		userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Role, &m.Text, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}
