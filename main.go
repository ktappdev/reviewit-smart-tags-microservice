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

func validateConfig() error {
	requiredVars := []string{"DATABASE_URL", "APP_API", "OPEN_ROUTER_API_KEY"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("required environment variable %s is not set", v)
		}
	}
	return nil
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
	loadEnv()
	
	if err := validateConfig(); err != nil {
		log.Fatalf("Configuration validation failed: %v", err)
	}
	
	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://reviewit.gy,http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Cache-Control, Authorization, X-Requested-With",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Content-Type",
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
