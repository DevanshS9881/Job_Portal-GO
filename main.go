package main
import(
	"github.com/gofiber/fiber/v2"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"log"
	//"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/middlewares"
	"github.com/DevanshS9881/Job_Portal-GO/handler"
)
func main(){
	app:=fiber.New()
	jwt:=middlewares.AuthMiddle(config.Secret)
	config.GoogleConfig()
	app.Post("/login",handler.Login)
	app.Get("/protected",jwt,handler.Protected)
	app.Get("/google_login", handler.GoogleLogin)
    app.Get("/google_callback", handler.GoogleCallback)
	app.Listen(":8080")

}