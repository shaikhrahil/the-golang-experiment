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

func New(r *fiber.Router, db *gorm.DB, conf rest.Configuration, logger *log.Logger) {
	repos := repositories{
		Auth:     NewRepository(db, logger),
		Accounts: accounts.NewRepository(db, logger),
	}
	NewController(r, conf, logger, repos.Auth, repos.Accounts)
	(*r).Use(GetMiddleware(conf))
	accounts.NewController(r, conf, logger, repos.Accounts)
}
