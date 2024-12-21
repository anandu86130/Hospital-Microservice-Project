package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will hasj the password and return the hash password
func HashPassword(pasword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

// CheckPassword function will check the provide password with userPassword
func CheckPassword(providedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	return err == nil
}
