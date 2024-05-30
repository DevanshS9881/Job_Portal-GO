package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	result := database.Db.Create(&newUser)
	//database.Db.Create(&newUser)
	if result.Error != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": result.Error,
		})
		return result.Error
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    newUser,
		"success": true,
		"message": "Successfully registered",
	})
	return nil
}
func UpdateProfileEmployee(c *fiber.Ctx) error {
	var newUser models.Employee
	userID := database.Convert(c.Params("id"))
	fmt.Println(userID)
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	newUser.UserID = userID
	 var existingUser models.User
	 
	 //checking whether the user exists or not
	 //if exits then update the record otherwise create it
	 
	if err:=database.Db.First(&existingUser,"id=?",userID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{"Error":err.Error})
	}
    database.Db.First(&existingUser,userID)
	var existingEmployee models.Employee
	if err:=database.Db.First(&existingEmployee,"user_id=?",userID).Error;err!=nil{
		result := database.Db.Create(&newUser)
		if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "data":    nil,
                "success": false,
                "message": result.Error.Error(),
            })
        }
	}else{
		result := database.Db.Model(&existingEmployee).Updates(newUser)
        if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "data":    nil,
                "success": false,
                "message": result.Error.Error(),
            })
        }
	}
	
	
		return c.Status(200).JSON(&fiber.Map{
		"data":    newUser,
		"success": true,
		"message": "Successfully Updated",
	})

}

func UpdateProfileEmployer(c *fiber.Ctx) error {
	var newUser models.Employer
	userID := database.Convert(c.Params("id"))
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"Error": err.Error(),
		})
	}
	newUser.UserID = userID
	var existingUser models.User
	if err:=database.Db.First(&existingUser,"id=?",userID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{ "Error":"User Not Found"})
	}
	var existingEmployer models.Employer
	if err:=database.Db.First(&existingEmployer,"user_id=?",userID).Error;err!=nil{
		result := database.Db.Create(&newUser)
		if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "data":    nil,
                "success": false,
                "message": result.Error.Error(),
            })
        }
	}else{
		result := database.Db.Model(&existingEmployer).Updates(newUser)
        if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "data":    nil,
                "success": false,
                "message": result.Error.Error(),
            })
        }
	}
	
	
		return c.Status(200).JSON(&fiber.Map{
		"data":    newUser,
		"success": true,
		"message": "Successfully Updated",
	})
	

}

func ShowProfile(c *fiber.Ctx) error {
	userID := database.Convert(c.Params("id"))
	// if err!=nil{
	// 	c.Status(400).JSON(fiber.Map{ "Error":"Invalid User ID"})
	// }
	var user models.User
	result := database.Db.Preload("Employee").Preload("Employee").First(&user, userID)
	if result.Error != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "No record exists",
		})
		return result.Error
	}
	c.Status(200).JSON(&fiber.Map{
		"data":    user,
		"success": true,
		"message": "Successfully Fetched",
	})
	return nil
}
