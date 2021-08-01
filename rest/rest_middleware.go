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
func GetTeam(c *fiber.Ctx) []uint {
	var teams []uint
	t := c.Locals("Teams").([]interface{})
	for _, v := range t {
		floatT := v.(float64)
		teams = append(teams, uint(floatT))
	}
	return teams
}

// TokenPayload defines the payload for the token
type TokenPayload struct {
	UserID uint
	TeamID uint
}
