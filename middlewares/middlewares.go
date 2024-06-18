package middlewares

import (
	"time"

	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
	//jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v4"
)
func AuthMiddle(secret string) fiber.Handler{
	// return jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(secret),
	// })
	// }

	jwtKey := []byte(secret)

	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("token")
		if accessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Missing or malformed JWT",
			})
		}

		claims := new(models.Claims)
		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error":   true,
						"message": "Malformed token",
					})
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error":   true,
						"message": "Token expired",
					})
				} else {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error":   true,
						"message": "Invalid token",
					})
				}
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": "Invalid token",
				})
			}
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid token",
			})
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Token expired",
			})
		}

		// Token is valid, proceed with the request
		return c.Next()
	}
}