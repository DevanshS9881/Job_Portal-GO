package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func Review(c *fiber.Ctx) error{
	Emid:=database.Convert(c.Params("Employer_id"))
	JobID:=database.Convert(c.Params("job_id"))
	var job models.Jobs
	var existingEmployer models.Employer
	if err:=database.Db.First(&existingEmployer,"user_id=?",Emid).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "message":"Invalid Employer ID"     ,            
		})
	}
	var existingEmployee models.Employee
    if err :=database.Db.Preload("Application").First(&job, "id = ? AND employer_id = ?", JobID, existingEmployer.ID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
    }

	if err :=database.Db.Preload("Application").First(&existingEmployee, "id = ?",job.Application[0].ID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
    }
    




    return c.Status(fiber.StatusOK).JSON(fiber.Map{"applications": job.Application})

}

// func Shortlist( c *fiber.Ctx) error{

