package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Generate generates the jwt token based on payload
func generate(payload jwt.MapClaims, expiresIn string, secret string) string {
	v, err := time.ParseDuration(expiresIn)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	payload["exp"] = time.Now().Add(v).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}

	return token
}

// Verify verifies the jwt token against the secret
func verify(token string, secret string) (map[string]interface{}, error) {
	parsed, err := parse(token, secret)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	u, ok := claims["UserID"]
	if !ok || u == "" {
		return nil, errors.New("something went wrong")
	}

	return claims, nil
}

func parse(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func (r *Repository) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
