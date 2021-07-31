package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

func New(r *fiber.Router, db *gorm.DB, config rest.Configuration, logger *log.Logger) {
	authRepo := auth.NewRepository(db, logger).SetConfig(auth.Config{
		Claims: NewClaims,
	})
	accountsRepo := accounts.NewRepository(db, logger)
	todoRepo := NewRepository(db, logger)

	auth.NewController(r, config, logger, authRepo, accountsRepo)
	(*r).Use(auth.GetMiddleware(config))
	accounts.NewController(r, config, logger, accountsRepo)
	NewController(r, config, logger, authRepo, accountsRepo, todoRepo)
}
