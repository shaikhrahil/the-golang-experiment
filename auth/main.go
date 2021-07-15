package main

import (
	"log"
	accounts "the-golang-experiment/accounts/lib"
	auth "the-golang-experiment/auth/lib"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	app := fiber.New()

	dsn := "auth-service:huhu@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}
	versioned := app.Group("/api/v1")
	logger := log.Default()
	logger.Println("Connected to DB")

	if err := db.AutoMigrate(&accounts.User{}); err != nil {
		log.Fatalln("Unable to migrate DB")
	}

	logger.Println("DB migrated")
	accounts.Controller(&versioned, db, logger)
	auth.Controller(&versioned, db, logger)

	app.Listen(":8000")
}