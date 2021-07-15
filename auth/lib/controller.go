package auth

import (
	"log"
	"the-golang-experiment/rest"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type controller struct {
	rest.Controller
}

func Controller(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	router := *r
	h := controller{
		Controller: rest.Controller{
			DB:     db,
			Logger: logger,
		},
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
