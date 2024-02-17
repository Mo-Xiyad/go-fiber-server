package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID        int    `json:"id"`
	Name     string `json:"name"`
	Age    string `json:"age"`
	Paid bool `json:"paid"`
}

func main (){
	app:= fiber.New()
	app.Use(logger.New())
	
	users := []User{}

  app.Get("/health-check", func(c *fiber.Ctx) error {
        return c.SendString("Server running!")
    })

		app.Post("/api/user", func(c *fiber.Ctx) error {
		user := &User{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		user.ID = len(users) + 1
		users = append(users, *user)
		return c.JSON(users)
		})


	app.Patch("/api/user/:id/paid", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Invalid ID provided: %v", err)
		return c.Status(400).SendString("Invalid ID")
	}
	
	found := false
	for index, user := range users {
		if user.ID == id {
			users[index].Paid = true
			found = true
			break
		}
	}
	
	if found {
		log.Printf("Updated user with ID: %d", id)
	} else {
		log.Printf("User with ID: %d not found", id)
	}
	
	return c.JSON(users)
})


app.Get("/api/user", func(c *fiber.Ctx) error {
	return c.JSON(users)
})

log.Fatal(app.Listen(":8080"))

}
