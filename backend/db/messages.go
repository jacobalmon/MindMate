package db

import (
	"context"
	"log"
)

func InsertMessage(userID int, role, text string) error {
	_, err := Pool.Exec(context.Background(),
		"INSRT INTO messages (user_id, role, text) VALUES ($1, $2, $3)", userID, role, text)
	if err != nil {
		log.Println("InsertMessage error", err)
	}
	return err
}

func GetLastMessages(userID, limit int) ([]string, error) {
	rows, err := Pool.Query(context.Background(),
		"SELECT text FROM messages WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2", userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		messages = append(messages, text)
	}
	return messages, nil
}
