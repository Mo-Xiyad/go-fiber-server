package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main() {
	app := fiber.New()

	// Middleware (Logger)
	app.Use(logger.New())

	// By doing this im simplifying the main.go file
	// ex: ProductRoutes(app)
	UserRoutes(app)

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server running!")
	})

	log.Fatal(app.Listen(":8080"))
}
