package routes

import (
	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/controllers"
	"github.com/DevanshS9881/Job_Portal-GO/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App){
	jwt:=middlewares.AuthMiddle(config.Secret)
	app.Post("/register",controllers.Register)
	app.Post("/updateProfileEmployee/:id",jwt,controllers.UpdateProfileEmployee)
	app.Post("/updateProfileEmployer/:id",jwt,controllers.UpdateProfileEmployer)
	app.Get("/getProfile/:id",jwt,controllers.ShowProfile)
	app.Delete("/deleteUser/:id",jwt,controllers.DeleteUser)
	app.Post("/addJob/:id",jwt,controllers.CreateJob)
	app.Get("/showJob/:id",jwt,controllers.GetJob)
	app.Delete("/deleteJob/:id/:Employer_id",jwt,controllers.DeleteJob)
	app.Post("/role",jwt,controllers.Role)
	app.Post("/apply",jwt,controllers.Apply)
	app.Get("/review/:Employer_id/:job_id",jwt,controllers.Review)
	app.Put("/updateJob/:id/:Employer_id",jwt,controllers.UpdateJob)

	//app.Post("/addJob/:id",con)
 }