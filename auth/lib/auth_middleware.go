package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

func GetMiddleware(config rest.Configuration) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
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
		user, err := verify(chunks[1], config.AUTH.JWT_SECRET)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		for k, v := range user {
			c.Locals(k, v)
		}
		// c.Locals("USER", user.UserID)
		// c.Locals("TEAM", user.TeamID)

		return c.Next()
	}

}
