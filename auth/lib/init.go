package auth

import (
	"log"

	"github.com/gofiber/fiber"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"gorm.io/gorm"
)

type repositories struct {
	Auth     Repository
	Accounts accounts.Repository
}

func New(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	repos := repositories{
		Auth:     *NewRepository(db, logger),
		Accounts: *accounts.NewRepository(db, logger),
	}
	NewController(r, logger, repos)
}
