package middlewares
import(
	"github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
)
func AuthMiddle(secret string) fiber.Handler{
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
	}
