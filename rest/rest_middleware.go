package rest

import (
	"github.com/gofiber/fiber/v2"
)

// GetUser is helper function for getting authenticated user's id
func GetUser(c *fiber.Ctx) uint {
	id := c.Locals("UserID").(float64)
	return uint(id)
}

// GetUser is helper function for getting authenticated user's id
func GetTeam(c *fiber.Ctx) []float64 {
	teams := c.Locals("Teams").([]float64)
	return teams
}

// TokenPayload defines the payload for the token
type TokenPayload struct {
	UserID uint
	TeamID uint
}
