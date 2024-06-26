package code

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(name *string) map[string]interface{} {

	payload := jwt.MapClaims{
		"sub":  "1234567890",
		"name": name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	payloadReturn := make(map[string]interface{})
	payloadReturn["token"] = tokenString

	return payloadReturn
}

func ValidateToken(tokenString *string) bool {
	token, _ := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if token.Valid {
		return token.Claims.Valid() == nil
	}

	return false
}
