package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func CreateJob(c *fiber.Ctx) error{
	newJob:=new(models.Jobs)
	if err:=c.BodyParser(newJob);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	result := database.Db.Create(&newJob)
if result.Error!=nil{
	c.Status(400).JSON(&fiber.Map{
		"data":    nil,
		"success": false,
		"message": "Cannot add new Job",
	})
	return result.Error
}
    c.Status(200).JSON(&fiber.Map{
		"data":    newJob,
		"success": true,
		"message": "Successfully added a new job",
	})
	return nil
}

func GetJob(c *fiber.Ctx) error{
			id:=database.Convert(c.Params("id"))
			job:=new(models.Jobs)
			result:=database.Db.First(&job,id)
			if result.Error!=nil{
				c.Status(400).JSON(&fiber.Map{
					"data":nil,
					"success":false,
					"message":"No record exists",
				})
				return result.Error
			}
			return c.Status(400).JSON(&fiber.Map{
				"data":job,
				"success":true,
				"message":"Successfully fetched the data",
			})
}

func DeleteJob(c *fiber.Ctx) error{
	id:=database.Convert(c.Params("id"))
	job:=new((models.Jobs))
	result:=database.Db.First(&job,id)
	if result.Error!=nil{
		c.Status(400).SendString("Invalid job id")
		return result.Error;
	}
	if err:=database.Db.Where("id=?",id).Delete(&models.Jobs{}).Error;err!=nil{
		c.Status(400).JSON(&fiber.Map{
			"data":nil,
			"success":false,
			"message":"No record exists",
		})
		return err
	}
	return c.Status(400).JSON(&fiber.Map{
		"data":job,
		"success":true,
		"message":"Successfully deleted the data",
	})
	
}

func GetJobByProfile(c *fiber.Ctx) error{
	profile:=c.Params("profile")
	var job []models.Jobs

}

