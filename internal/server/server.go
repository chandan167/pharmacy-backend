package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chandan167/pharmacy-backend/internal/router"
	"github.com/chandan167/pharmacy-backend/pkg/helper"
	"github.com/chandan167/pharmacy-backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const IdealTimeout = time.Second * 5

func Server() *fiber.App {
	prefork := false
	if os.Getenv("GO_ENV") == "production" {
		prefork = true
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		IdleTimeout:  IdealTimeout,
		Prefork:      prefork,
	})
	// Middlewares
	app.Use(requestid.New()) // assign request ID
	app.Use(recover.New())   // catch panics safely
	// Initialize default config
	app.Use(pprof.New())

	router.SetUpRoute(app)
	return app
}

func Start(app *fiber.App) {
	go func() {
		log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	}()
	graceFullShutDown(app)
}

func NormalStart(app *fiber.App) {
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func graceFullShutDown(app *fiber.App) {
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "inter server error"
	var validation_error any = nil

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Error()
	}

	if e, ok := err.(*helper.AppError); ok {
		code = e.StatusCode
		message = e.Error()
		validation_error = e.ValidationError
	}

	reqId := ctx.Locals("requestid")

	logger.Logger.Error("ERROR HANDLER",
		"STATUS", code,
		"MESSAGE", message,
		"REQUEST_ID", reqId,
		"METHOD", ctx.Method(),
		"PATH", ctx.Path(),
		"VALIDATION_ERROR", validation_error,
	)

	return ctx.Status(code).JSON(fiber.Map{
		"success":          false,
		"status":           code,
		"message":          message,
		"request_id":       reqId,
		"method":           ctx.Method(),
		"path":             ctx.Path(),
		"validation_error": validation_error,
	})
}
