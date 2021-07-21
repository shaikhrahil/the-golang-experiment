package accounts

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type repositories struct {
	Accounts Repository
}

func New(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	Accounts := NewRepository(db, logger)
	NewController(r, logger, Accounts)
}
