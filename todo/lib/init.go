package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"gorm.io/gorm"
)

type repositories struct {
	Auth     auth.Repository
	Accounts accounts.Repository
	Todo     Repository
}

func New(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	repos := repositories{
		Auth:     auth.NewRepository(db, logger),
		Accounts: accounts.NewRepository(db, logger),
		Todo:     NewRepository(db, logger),
	}
	protected := (*r).Use(auth.Middleware)
	accounts.NewController(&protected, logger, repos.Accounts)
	auth.NewController(r, logger, repos.Auth, repos.Accounts)
	NewController(&protected, logger, repos.Auth, repos.Accounts)
}