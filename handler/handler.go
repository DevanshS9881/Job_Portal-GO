package handler
import(
	"time"
    "github.com/gofiber/fiber/v2"
	jtoken"github.com/golang-jwt/jwt/v4"
	"github.com/DevanshS9881/Job_Portal-G0/src/github.com/DevanshS9881/Job_Portal-GO/internal/config"
	"github.com/DevanshS9881/Job_Portal-G0/internal/models"
	"github.com/DevanshS9881/Job_Portal-G0/internal/repository"
)
func Login(c *fiber.Ctx) error{
	loginRequest:=new(models.LoginRequest)
	if err:=c.BodyParser(loginRequest);err!=nil{
		return c.Status(fiber.StatusBadRequest).JS
	}
}