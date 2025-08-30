package db

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Hashing password failed", err)
		return err
	}

	_, err = Pool.Exec(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)", email, string(hashedPassword))
	if err != nil {
		log.Println("CreateUser Error", err)
	}
	return err
}

func CheckUserCredentials(email, password string) (bool, error) {
	var hashedPassword string
	err := Pool.QueryRow(context.Background(),
		"SELECT password_hash FROM users WHERE email=$1", email).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}
