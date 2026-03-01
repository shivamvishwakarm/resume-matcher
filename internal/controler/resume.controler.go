package controler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetAllResume(c *fiber.Ctx) error {

	user, ok := c.Locals("user").(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	userEmail, _ := user["email"].(string)
	log.Println("Getting all resumes for user:", userEmail)

	// Query MongoDB for all resumes belonging to this user
	coll := config.Db.Collection("resumes")
	filter := bson.M{"user_email": userEmail}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error fetching resumes from MongoDB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch resumes",
		})
	}
	defer cursor.Close(context.TODO())

	var resumes []models.Resume
	if err := cursor.All(context.TODO(), &resumes); err != nil {
		log.Println("Error decoding resumes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode resumes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total":   len(resumes),
		"resumes": resumes,
	})
}
