package config

import (
	"api-gateway/internal/http/token"
	"database/sql"
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func IsValidPassword(password string) error {
	if len(password) < 5 {
		return errors.New("password must be at least 5 characters long")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidUUID(id string) error {
	if err := uuid.Validate(id); err != nil {
		return errors.New("invalid user uuid")
	}
	return nil
}

func CheckRowsAffected(res sql.Result, object string) error {
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New(object + " not found")
	}
	return nil
}

func GetClaims(c *gin.Context) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header required")
	}

	claims, err := token.ExtractClaim(authHeader)
	if err != nil {
		return nil, errors.New("invalid token claims" + err.Error())
	}
	return claims, nil
}

func GetRefreshToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header required")
	}

	valid, err := token.ValidateToken(authHeader)
	if err != nil || !valid {
		return "", errors.New("invalid token" + err.Error())
	}

	return authHeader, nil
}
