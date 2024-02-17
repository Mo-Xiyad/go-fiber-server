// user.go

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Paid  bool   `json:"paid"`
}
// simulates db
var users = []User{}

// add all user routes here
func UserRoutes(app *fiber.App) {
	app.Post("/api/user", createUser)
	app.Patch("/api/user/:id/paid", updateUserPaidStatus)
	app.Get("/api/user", getUsers)
}

func createUser(c *fiber.Ctx) error {
	user := &User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user.ID = len(users) + 1
	users = append(users, *user)
	return c.JSON(users)
}

func updateUserPaidStatus(c *fiber.Ctx) error {
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
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}
