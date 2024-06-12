package controllers

import (
	"fmt"
	"time"

	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/DevanshS9881/Job_Portal-GO/hashPassword"

)

func Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	newUser.Password,_=hashpassword.HashPassword(newUser.Password)
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
	//return nil
	day:=time.Hour*24;
	claims:=jtoken.MapClaims{
		"ID": newUser.ID,
		"email":newUser.Email,
		"expi":time.Now().Add(day*1).Unix(),
	}
	token:=jtoken.NewWithClaims(jtoken.SigningMethodHS256,claims)
	t,err:=token.SignedString([]byte(config.Secret))
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error":err.Error(),})
	}
	return c.JSON(models.LoginResponse{
		Token:t,
	})
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
	//newUser.Password,_=hashpassword.HashPassword(newUser.Password)
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

func DeleteUser(c *fiber.Ctx) error{
	id:=c.Params("id")
	user:=new((models.User))
	result:=database.Db.Preload("Employee").Preload("Employer").First(&user,id)
	if result.Error!=nil{
		c.Status(400).SendString("Invalid user id")
		return result.Error;
	}
	if err:=database.Db.Where("id=?",id).Delete(&models.User{}).Error;err!=nil{
		c.Status(400).JSON(&fiber.Map{
			"data":nil,
			"success":false,
			"message":"No record exists",
		})
		return err
	}
	if user.Employee.ID!=0{
		if err:=database.Db.Where("user_id=?",id).Delete(&models.Employee{}).Error;err!=nil{
			c.Status(400).JSON(&fiber.Map{
				"data":nil,
				"success":false,
				"message":"No employee record exists",
			})
			return err
		}
	}else{
		if err:=database.Db.Where("user_id=?",id).Delete(&models.Employer{}).Error;err!=nil{
			c.Status(400).JSON(&fiber.Map{
				"data":nil,
				"success":false,
				"message":"No employer record exists",
			})
			return err
		}
	}
	return c.Status(400).JSON(&fiber.Map{
		"data":user,
		"success":true,
		"message":"Successfully deleted the data",
	})
	
}

func Role(c *fiber.Ctx) error{
	var currUser models.User
	var currRole models.Roles
	if err:=c.BodyParser(&currRole);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
    user:=c.Locals("user").(*jtoken.Token)
	claims:=user.Claims.(jtoken.MapClaims)
	idF6:=claims["ID"].(float64)
	id:=uint(idF6)
	if result:=database.Db.First(&currUser,id).Error;result!=nil{
		return c.Status(400).JSON(&fiber.Map{
			"message":"User does not exist",
		})
	}
	fmt.Println(currRole.Role)
	currUser.Role=currRole.Role
	fmt.Println(currUser.Role)
	database.Db.Save(&currUser)
	return nil
}
