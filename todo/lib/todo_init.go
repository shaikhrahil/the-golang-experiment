package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type repositories struct {
	Auth     auth.Repository
	Accounts accounts.Repository
	Todo     Repository
}

func New(r *fiber.Router, db *gorm.DB, config rest.Configuration, logger *log.Logger) {
	repos := repositories{
		Auth:     auth.NewRepository(db, logger),
		Accounts: accounts.NewRepository(db, logger),
		Todo:     NewRepository(db, logger),
	}
	auth.NewController(r, config, logger, repos.Auth, repos.Accounts)
	(*r).Use(auth.GetMiddleware(config))
	accounts.NewController(r, config, logger, repos.Accounts)
	NewController(r, config, logger, repos.Auth, repos.Accounts, repos.Todo)
}
