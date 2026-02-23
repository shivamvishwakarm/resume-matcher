package middleware

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Get("Authorization")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		parts := strings.Split(token, ".")
		if len(parts) != 3 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token format",
			})
		}

		payload, err := jwt.DecodeSegment(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var claims map[string]interface{}
		if err := json.Unmarshal(payload, &claims); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token payload",
			})
		}

		// Store the parsed claims in context locals under the key "user"
		c.Locals("user", claims)

		return c.Next()
	}
}
