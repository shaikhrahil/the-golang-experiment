package main

import (
	"fmt"
	"log"

	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	app := fiber.New()

	config := rest.GetConfig("auth")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, config.DB.USERNAME, config.DB.PASSWORD, config.DB.HOST, config.DB.PORT, config.DB.NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}
	versioned := app.Group(fmt.Sprintf("/api/%s", config.APP.VERSION))
	logger := log.Default()
	logger.Println("Connected to DB")

	if err := db.AutoMigrate(&accounts.User{}); err != nil {
		log.Fatalln("Unable to migrate DB")
	}

	logger.Println("DB migrated")

	auth.New(&versioned, db, config, logger)

	logger.Fatalln(app.Listen(fmt.Sprintf(`:%s`, config.APP.PORT)))
}
