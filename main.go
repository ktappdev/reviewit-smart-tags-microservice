package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Using system environment variables.")
	}
}

func validateConfig() error {
	if os.Getenv("OPEN_ROUTER_API_KEY") == "" {
		return fmt.Errorf("OPEN_ROUTER_API_KEY environment variable is not set")
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

	app.Post("/gen", getTagsFromDescription)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Fatal(app.Listen(":" + port))
}
