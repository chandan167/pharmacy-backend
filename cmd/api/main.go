package main

import (
	"flag"
	"log"
	"os"

	"github.com/chandan167/pharmacy-backend/pkg/logger"
	"github.com/chandan167/pharmacy-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	router.SetUpRoute(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func init() {
	envFile := flag.String("env", ".env", "Path to .env file")
	flag.Parse()
	if err := godotenv.Load(*envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", *envFile, err)
	}
	logger.Init(os.Getenv("GO_ENV"), os.Getenv("LOG_FILE"))
}
