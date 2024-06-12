package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func Apply(c *fiber.Ctx) error{
	newAppl:=new(models.Application)
	if err := c.BodyParser(newAppl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	result := database.Db.Create(&newAppl)
	
	if result.Error != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": result.Error,
		})
		return result.Error
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    newAppl,
		"success": true,
		"message": "Successfully registered",
	})
	return nil
}


