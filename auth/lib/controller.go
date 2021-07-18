package auth

import (
	"log"

	"github.com/gofiber/fiber"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
)

type controller struct {
	logger *log.Logger
	repos  repositories
}

func NewController(r *fiber.Router, logger *log.Logger, repos map[string]Repository) {
	router := *r
	myRepos := repositories{
		Auth:     repos["Auth"],
		Accounts: accounts.Repository{},
	}
	h := controller{
		logger: logger,
		repos:  myRepos,
	}
	authRoutes := router.Group("/auth")
	authRoutes.Post("/login", h.Login)
	authRoutes.Post("/logout", h.Logout)
	authRoutes.Post("/token/refresh", h.RefreshToken)
	authRoutes.Post("/token/validate", h.ValidateToken)
}

func (u *controller) Login(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"awesome": "hello world !!",
	})
}

func (u *controller) Logout(c *fiber.Ctx) {

}

func (u *controller) RefreshToken(c *fiber.Ctx) {

}

func (u *controller) ValidateToken(c *fiber.Ctx) {

}
