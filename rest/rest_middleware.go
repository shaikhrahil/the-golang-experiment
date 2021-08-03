package rest

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/gofiber/fiber/v2"
)

// GetUser is helper function for getting authenticated user's id
func GetUser(c *fiber.Ctx) uint {
	id := c.Locals("UserID").(float64)
	return uint(id)
}

// GetUser is helper function for getting authenticated user's id
func GetTeams(c *fiber.Ctx) mapset.Set {
	var teams mapset.Set
	t := c.Locals("Teams").([]interface{})
	for _, v := range t {
		// floatT :=
		// teams = append(teams, uint(floatT))
		teams.Add(v.(uint))
	}
	return teams
}

// TokenPayload defines the payload for the token
type TokenPayload struct {
	UserID uint
	TeamID uint
}
