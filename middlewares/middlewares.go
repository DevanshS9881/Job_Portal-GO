package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jwt "github.com/golang-jwt/jwt/v4"
)

func AuthMiddle(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			
			// Check token expiration
			expiration := time.Unix(int64(claims["expi"].(float64)), 0)
			if time.Now().After(expiration) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token has expired",
				})
			}
			
			// Continue processing if token is valid
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	})
}
