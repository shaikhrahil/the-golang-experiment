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
	migrate(db)
	repos := repositories{
		Auth:     auth.NewRepository(db, logger),
		Accounts: accounts.NewRepository(db, logger),
		Todo:     NewRepository(db, logger),
	}
	auth.NewController(r, logger, repos.Auth, repos.Accounts)
	(*r).Use(auth.Middleware)
	accounts.NewController(r, logger, repos.Accounts)
	NewController(r, logger, repos.Auth, repos.Accounts, repos.Todo)
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &Todo{}, UserTodo{}); err != nil {
		log.Fatalln("Unable to migrate DB")
	}

	// if err := db.SetupJoinTable(&User{}, "Todos", &UserTodo{}); err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// if err := db.SetupJoinTable(&Todo{}, "UserID", &UserTodo{}); err != nil {
	// 	log.Fatalln(err.Error())
	// }

}
