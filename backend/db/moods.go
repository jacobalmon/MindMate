package db

import (
	"context"
	"log"
)

func InsertMood(userID int, note string, mood_score int) error {
	_, err := Pool.Exec(context.Background(),
		"INSERT INTO moods (user_id, note, mood_score) VALUES ($1, $2, $3)", userID, note, mood_score)
	if err != nil {
		log.Println("InsertMood error", err)
	}
	return err
}

func GetMoodHistory(userID int) ([]string, error) {
	rows, err := Pool.Query(context.Background(),
		"SELECT mood FROM moods WHERE user_id=$1 ORDER BY created_at DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moods []string
	for rows.Next() {
		var mood string
		if err := rows.Scan(&mood); err != nil {
			return nil, err
		}
		moods = append(moods, mood)
	}
	return moods, nil
}
