package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(value string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))
	return err == nil
}
