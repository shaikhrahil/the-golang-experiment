package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
)

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}

// TokenPayload defines the payload for the token
type TokenPayload struct {
	ID uint
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload, expiresIn string, secret string) string {
	v, err := time.ParseDuration(expiresIn)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}

	return token
}

// Verify verifies the jwt token against the secret
func Verify(token string, secret string) (*TokenPayload, error) {
	parsed, err := parse(token, secret)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID: uint(id),
	}, nil
}

func parse(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
