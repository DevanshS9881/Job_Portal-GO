package main

import (
	"github.com/gofiber/fiber/v2"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"log"
	//"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/handler"
	"github.com/DevanshS9881/Job_Portal-GO/middlewares"
	"github.com/DevanshS9881/Job_Portal-GO/routes"
)
func main(){
	app:=fiber.New()
	dbEr:=database.InitDB()
	if dbEr!=nil{
		panic("Connection failed to databse")
	}
	jwt:=middlewares.AuthMiddle(config.Secret)
	config.GoogleConfig()
	app.Post("/login",handler.Login)
	app.Get("/protected",jwt,handler.Protected)
	app.Get("/google_login", handler.GoogleLogin)
    app.Get("/google_callback", handler.GoogleCallback)
	routes.SetRoutes(app)
	app.Listen(":8080")


}

