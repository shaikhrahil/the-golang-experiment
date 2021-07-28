package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/config"
	todo "github.com/shaikhrahil/the-golang-experiment/todo/lib"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()
	config := config.GetConfig("todo")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	migrate(db)
	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}
	versioned := app.Group("/api/v1")
	logger := log.Default()
	logger.Println("Connected to DB")

	logger.Println("DB migrated")

	todo.New(&versioned, db, logger)

	app.Listen(":8000")
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&user.User{}, &todo.Todo{}); err != nil {
		log.Fatalln("Unable to migrate DB")
	}

	// if err := db.SetupJoinTable(&User{}, "Todos", &UserTodo{}); err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// if err := db.SetupJoinTable(&Todo{}, "UserID", &UserTodo{}); err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// if err := db.SetupJoinTable(&Todo{}, "Todos", &User{}); err != nil {
	// 	log.Fatalln(err.Error())
	// }

}
