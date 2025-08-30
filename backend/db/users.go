package db

import (
	"context"
	"log"
)

func CreateUser(username, hashedPassword string) error {
	_, err := Pool.Exec(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		log.Println("CreateUser Error", err)
	}
	return err
}

func CheckUserCredentials(username, hashedPassword string) (bool, error) {
	var exists bool
	err := Pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1 AND password=$2)", username, hashedPassword).Scan(&exists)
	return exists, err
}
