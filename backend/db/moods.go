package db

import (
	"context"
	"log"
	"time"
)

type Mood struct {
	ID        int       `json:"id"`
	Note      string    `json:"note"`
	Score     int       `json:"mood_score"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertMood(userID int, note string, mood_score int) error {
	_, err := Pool.Exec(context.Background(),
		"INSERT INTO moods (user_id, note, mood_score) VALUES ($1, $2, $3)", userID, note, mood_score)
	if err != nil {
		log.Println("InsertMood error", err)
	}
	return err
}

func GetMoodHistory(userID, limit int) ([]Mood, error) {
	rows, err := Pool.Query(context.Background(),
		"SELECT id, note, mood_score, created_at FROM moods WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2",
		userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moods []Mood
	for rows.Next() {
		var m Mood
		if err := rows.Scan(&m.ID, &m.Note, &m.Score, &m.CreatedAt); err != nil {
			return nil, err
		}
		moods = append(moods, m)
	}

	for i, j := 0, len(moods)-1; i < j; i, j = i+1, j-1 {
		moods[i], moods[j] = moods[j], moods[i]
	}

	return moods, nil
}
