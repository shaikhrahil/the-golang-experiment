package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team_user"
	"gorm.io/gorm"
)

func New(r *fiber.Router, db *gorm.DB, config rest.Configuration, logger *log.Logger) {
	teamUsersRepo := team_user.NewRepository(db, logger)
	authRepo := auth.NewRepository(db, logger).SetConfig(auth.Config{
		Claims: NewClaims(teamUsersRepo, logger),
	})
	accountsRepo := accounts.NewRepository(db, logger)
	teamsRepo := team.NewRepository(db, logger)
	todoRepo := NewRepository(db, logger)

	auth.NewController(r, config, logger, authRepo, accountsRepo)
	(*r).Use(auth.GetMiddleware(config))
	team.NewController(r, config, logger, teamsRepo)
	team_user.NewController(r, config, logger, teamUsersRepo)
	accounts.NewController(r, config, logger, accountsRepo)
	NewController(r, config, logger, authRepo, accountsRepo, todoRepo)
}
