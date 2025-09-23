package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/chandan167/pharmacy-backend/pkg/logger"
	"github.com/chandan167/pharmacy-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	// Middlewares
	app.Use(requestid.New()) // assign request ID
	app.Use(recover.New())   // catch panics safely
	// Initialize default config
	app.Use(pprof.New())

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

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "inter server error"
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	reqId := ctx.Locals("requestid")

	logger.Logger.Error("ERROR HANDLER",
		"STATUS", code,
		"MESSAGE", message,
		"REQUEST_ID", reqId,
		"METHOD", ctx.Method(),
		"PATH", ctx.Path(),
	)

	return ctx.Status(code).JSON(fiber.Map{
		"success":    false,
		"status":     code,
		"message":    message,
		"request_id": reqId,
		"method":     ctx.Method(),
		"path":       ctx.Path(),
	})
}
