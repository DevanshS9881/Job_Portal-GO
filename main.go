package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

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
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://127.0.0.1:3002",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))
	config.GoogleConfig()
	app.Static("/", "/frontend")
	app.Post("/login",handler.Login)
	app.Get("/protected",jwt,handler.Protected)
	app.Get("/google_login", handler.GoogleLogin)
    app.Get("/google_callback",handler.GoogleCallback)
	routes.SetRoutes(app)
	err := app.Listen(":8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}

