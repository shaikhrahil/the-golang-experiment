package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type Configuration struct {
	rest.Configuration
}

type repositories struct {
	Auth     Repository
	Accounts accounts.Repository
}

func New(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	repos := repositories{
		Auth:     NewRepository(db, logger),
		Accounts: accounts.NewRepository(db, logger),
	}
	NewController(r, logger, repos.Auth, repos.Accounts)
	(*r).Use(Middleware)
	accounts.NewController(r, logger, repos.Accounts)
}
