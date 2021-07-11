package main

import (
	"the-golang-experiment/auth/user"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	versioned := app.Group("v1")

	user.AddHandlers(versioned)

	app.Listen(":8000")
}
