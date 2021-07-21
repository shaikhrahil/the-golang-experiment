package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Package struct {
	Repository func(db *gorm.DB, logger *log.Logger) Repository
	Controller func(r *fiber.Router, logger *log.Logger) Controller
}

type Repository struct {
	DB     *gorm.DB
	Logger *log.Logger
}

type Controller struct {
	Logger *log.Logger
}
