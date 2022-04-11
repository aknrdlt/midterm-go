package app

import (
	"log"
	"os"
	"os/signal"
	"rd/internal/api"
	"rd/internal/api/handlers"
	"rd/internal/db"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

const idleTimeout = 10 * time.Second

// Run initializes whole application
func Run() {
	_app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			return c.Status(code).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		},
		DisableStartupMessage: true,
	})

	_db := db.DB()
	defer _db.Close()

	_dbm := db.NewDataBaseLayers(_db)
	_handler := handlers.NewHandler(_dbm)

	api.Routes(_app, _handler)
	go func() {
		if err := _app.Listen(":" + "8080"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create chanel to segnify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c //This block the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = _app.Shutdown()

	log.Println("Running cleanup tasks...")
}
