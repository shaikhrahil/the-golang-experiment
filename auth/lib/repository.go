package auth

import (
	"log"
	"time"

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
func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration("2000000000")

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte("config.TOKENKEY"))

	if err != nil {
		panic(err)
	}

	return token
}
