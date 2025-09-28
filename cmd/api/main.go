package main

import (
	"flag"
	"log"
	"os"

	"github.com/chandan167/pharmacy-backend/internal/server"
	"github.com/chandan167/pharmacy-backend/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := server.Server()
	if os.Getenv("GO_ENV") == "development" {
		server.NormalStart(app)
	} else {
		server.Start(app)
	}

}

func init() {
	envFile := flag.String("env", ".env", "Path to .env file")
	flag.Parse()
	if err := godotenv.Load(*envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", *envFile, err)
	}
	logger.Init(os.Getenv("GO_ENV"), os.Getenv("LOG_FILE"))
}
