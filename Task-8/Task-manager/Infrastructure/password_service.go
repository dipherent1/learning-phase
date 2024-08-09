package infrastructure

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Internal server error")
	}

	hp := string(hashedPassword)

	return hp, nil
}

func CheckPassword(existingPassword string, loginPassword string) error {
	
	if bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(loginPassword)) != nil {
		return errors.New("invalid password")

	}

	return nil
}
