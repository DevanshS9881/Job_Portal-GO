package routes

import (
	"github.com/DevanshS9881/Job_Portal-GO/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App){
	app.Post("/register",controllers.Register)
	app.Post("/updateEmployee/:id",controllers.UpdateProfileEmployee)
	app.Get("/updateEmployer/:id",controllers.UpdateProfileEmployer)
 }