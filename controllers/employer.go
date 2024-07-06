package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func Review(c *fiber.Ctx) error {
    Emid := database.Convert(c.Params("Employer_id"))
    JobID := database.Convert(c.Params("job_id"))
    var job models.Jobs
    var existingEmployer models.Employer

    if err := database.Db.First(&existingEmployer, "user_id = ?", Emid).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "message": "Invalid Employer ID",
        })
    }

    // Preload Application and Employee data
    if err := database.Db.Preload("Application.Employee").First(&job, "id = ? AND employer_id = ?", JobID, existingEmployer.ID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"applications": job.Application})
}

func Accept( c *fiber.Ctx) error{
	apID:=database.Convert(c.Params("id"))
	var application models.Application
	if err := c.BodyParser(&application); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	var existingApplication models.Application
	if err := database.Db.First(&existingApplication, "id = ?",apID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "message": "Invalid Application ID",
        })
    }
	existingApplication.Review=application.Review
	database.Db.Save(&existingApplication)
     return nil
}

