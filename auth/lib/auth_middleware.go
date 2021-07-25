package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
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
