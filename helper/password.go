package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(pass string) (string, error)  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckHashPassword(pass, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
    return err == nil
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "Failed to get token", err
	}
	return tokenString, nil
}