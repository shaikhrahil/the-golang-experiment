package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	todo "github.com/shaikhrahil/the-golang-experiment/todo/lib"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()
	config := rest.GetConfig("todo")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, config.DB.USERNAME, config.DB.PASSWORD, config.DB.HOST, config.DB.PORT, config.DB.NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	migrate(db)
	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}
	versioned := app.Group(fmt.Sprintf("/api/%s", config.APP.VERSION))
	logger := log.Default()
	logger.Println("Connected to DB")

	logger.Println("DB migrated")

	todo.New(&versioned, db, config, logger)

	app.Listen(fmt.Sprintf(`:%s`, config.APP.PORT))
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&user.User{}, &todo.Todo{}); err != nil {
		log.Fatalln("Unable to migrate DB")
	}
}
