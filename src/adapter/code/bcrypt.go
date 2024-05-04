package code

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) string {
	pwd := *password
	pwdBytes := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}
