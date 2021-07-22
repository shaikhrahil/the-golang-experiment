package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Middleware(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")

	}

	// Spliting the header
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	// Verify the token which is in the chunks
	user, err := Verify(chunks[1])

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	c.Locals("USER", user.ID)

	return c.Next()
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// Getting ID, it's an interface{} so I need to cast it to uint
	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID: uint(id),
	}, nil
}

func parse(token string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("config.TOKENKEY"), nil
	})
}
