package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 3)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ValidatePassword(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
