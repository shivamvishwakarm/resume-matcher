package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"github.com/shivamvishwakarm/resume-matcher/internal/controler"
	"github.com/shivamvishwakarm/resume-matcher/internal/middleware"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
)

func main() {

	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/register", createuser)
	v1.Post("/login", login)

	v1.Get("/me", middleware.Auth(), controler.GetUser())

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("health")
	})

	log.Println("Server starting on port:8080")
	log.Fatal(app.Listen(":8080"))
}

func login(c *fiber.Ctx) error {

	var user config.LoginReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := config.LoginUser(user)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON((fiber.Map{
			"error": err.Error(),
		}))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    result.Token,
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(result)
}

func createuser(c *fiber.Ctx) error {
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
}
