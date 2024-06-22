package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	saltStr := os.Getenv("SALT")

	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
