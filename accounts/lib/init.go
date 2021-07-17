package accounts

import (
	"log"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func Init(r *fiber.Router, db *gorm.DB, logger *log.Logger) {
	Controller(r, db, logger)
	Repository(db, logger)
}
