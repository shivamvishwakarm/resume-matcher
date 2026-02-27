package controler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MeRes struct {
	Id        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

func GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		coll := config.Client.Database("resume-matcher").Collection("user")

		// Retrieve the user claims stored by the auth middleware
		claims, ok := c.Locals("user").(map[string]interface{})
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email not found in token",
			})
		}

		filter := bson.M{"email": email}

		var user MeRes
		err := coll.FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(user)
	}
}
