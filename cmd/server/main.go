package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		createdUser, err := config.CreateUser(user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		createdUser.Password = ""

		return c.Status(fiber.StatusCreated).JSON(createdUser)
	})

	log.Println("Server starting on port:8080")
	log.Fatal(app.Listen(":8080"))
}
