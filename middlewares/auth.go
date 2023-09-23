package middlewares

import (
	"elearning/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func CheckToken(tokenString string) (string, error) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return "", errors.New("token invalid")
	}

	if !token.Valid {
		return "", errors.New("expired")
	}

	return claims.Email, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]

		if header == nil {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			_, err := c.Writer.Write([]byte("unauthorized"))
			if err != nil {
				return
			}
			return
		}

		split := strings.Split(header[0], " ")
		if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			_, err := c.Writer.Write([]byte("bearer token format needed"))
			if err != nil {
				return
			}
			return
		}

		_, err := CheckToken(split[1])
		if err != nil {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			_, err := c.Writer.Write([]byte("token invalid"))
			if err != nil {
				return
			}
			return
		}
	}
}
