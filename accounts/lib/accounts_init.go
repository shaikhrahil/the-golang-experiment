package accounts

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

func New(r *fiber.Router, db *gorm.DB, conf rest.Configuration, logger *log.Logger) {
	Accounts := NewRepository(db, logger)
	NewController(r, conf, logger, Accounts)
}
