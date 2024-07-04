package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func Apply(c *fiber.Ctx) error{
	newAppl:=new(models.Application)
	employeeID := database.Convert(c.Params("Emid"))
	jobID:=database.Convert(c.Params("jobID"))
	if err := c.BodyParser(newAppl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	var existingEmployee models.Employee
	if err:=database.Db.First(&existingEmployee,"user_id=?",employeeID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "Message":"Invalid Employee"     ,            
		})
	}
	var existingJob models.Jobs
	if err:=database.Db.First(&existingJob,"id=?",jobID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "Message":"Invalid Job"     ,            
		})
	}
	newAppl.EmployeeID=existingEmployee.ID
	newAppl.JobsID=jobID
	result := database.Db.Create(&newAppl)
	
	if result.Error != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": result.Error,
		})
		return result.Error
	}
	existingJob.ApplicationsRecieved++;
    database.Db.Save(&existingJob)
	c.Status(200).JSON(&fiber.Map{
		"data":    newAppl,
		"success": true,
		"message": "Successfully registered",
	})
	return nil
}


