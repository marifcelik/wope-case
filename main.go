package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fiber-api/models"
	"fiber-api/routers"
)

//go:generate go run github.com/steebchen/prisma-client-go generate
func main() {
	godotenv.Load()
	log.Println("Starting server v2...")
	app := fiber.New()
	PORT := getPort("3000")

	DB_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	db.AutoMigrate(&models.Task{})

	app.Use(fiberlogger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hi fiber")
	})

	routers.InitNoteRouter(app, db)

	log.Fatal(app.Listen(":" + PORT))
}

func getPort(def string) string {
	port := os.Getenv("PORT")
	if port == "" {
		return def
	}
	return port
}
