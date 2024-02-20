package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Middleware (Logger)
	app.Use(logger.New())

	// creating a route group for /api
	api := app.Group("/api")

	// By doing this we can add more routes to the /api group
	// ex: ProductRoutes(app)
	UserRoutes(api) // add user routes to /api/user

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server running!")
	})

	log.Fatal(app.Listen(":8080"))
}
