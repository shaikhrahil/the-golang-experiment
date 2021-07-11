package rest

import (
	"log"

	"gorm.io/gorm"
)

type Controller struct {
	Logger *log.Logger
	DB     *gorm.DB
}
