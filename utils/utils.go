package utils

import (
	"elearning/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePassword(hashed string, current []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), current)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GenerateToken(email string) (*string, error) {
	expire := time.Now().Add(24 * time.Hour)

	claims := models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err == nil {
		return &tokenString, nil
	}

	return nil, errors.New(err.Error())
}
