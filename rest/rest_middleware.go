package rest

import "github.com/gofiber/fiber/v2"

// GetUser is helper function for getting authenticated user's id
func GetUser(c *fiber.Ctx) uint {
	id, _ := c.Locals("USER").(uint)
	return id
}
