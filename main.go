package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Using system environment variables.")
	}
}

func initDB() error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}

func main() {
	loadEnv() // Load .env file if it exists
	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{ // Apply CORS middleware
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000, https://reviewit.lugetech.com/", // Allow requests from localhost
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	app.Post("/gettags", getTags)
	app.Post("/regen", regenerateTags)
	app.Post("/gen", getTagsFromDescription)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Fatal(app.Listen(":" + port))
}
