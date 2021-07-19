package auth

import (
	"log"

	"github.com/gofiber/fiber"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
)

type controller struct {
	logger         *log.Logger
	authService    Repository
	accountService accounts.Repository
}

func NewController(r *fiber.Router, logger *log.Logger, authService Repository, accountsService accounts.Repository) {
	router := *r
	h := controller{
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
	}
	authRoutes := router.Group("/auth")
	authRoutes.Get("/login", h.Login)
	authRoutes.Post("/logout", h.Logout)
	authRoutes.Post("/token/refresh", h.RefreshToken)
	authRoutes.Post("/token/validate", h.ValidateToken)
}

func (u *controller) Login(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"accountsService": u.accountService.GetSomething(),
		"authService":     u.authService.GetSomething(),
	})
}

func (u *controller) Logout(c *fiber.Ctx) {

}

func (u *controller) RefreshToken(c *fiber.Ctx) {

}

func (u *controller) ValidateToken(c *fiber.Ctx) {

}
