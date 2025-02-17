package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	 bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	 return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	fmt.Printf("password: %s, hashedPassword: %s\n", password, hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
