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
	result := database.Db.Preload("Employee").Preload("Employer").First(&user, userID)
	if result.Error != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "No record exists",
		})
		
	}
	return c.Status(200).JSON(&fiber.Map{
		"data":    user,
		"success": true,
		"message": "Successfully Fetched",
	})
	
}

func DeleteUser(c *fiber.Ctx) error {
    // Parse the user ID from the URL parameters
    userID := database.Convert(c.Params("id"))

    // Convert userID to uint
    var user models.User
    if err := database.Db.First(&user, userID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    // Start a transaction
    tx := database.Db.Begin()
    if tx.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to start transaction",
        })
    }

    var existingEmployer models.Employer
    if err := database.Db.First(&existingEmployer, "user_id = ?", userID).Error; err == nil {
        // Delete related applications for each job
        var jobs []models.Jobs
        if err := tx.Where("employer_id = ?", existingEmployer.ID).Find(&jobs).Error; err == nil {
            for _, job := range jobs {
                if err := tx.Where("jobs_id = ?", job.ID).Delete(&models.Application{}).Error; err != nil {
                    tx.Rollback()
                    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                        "error": "Failed to delete applications",
                    })
                }
            }
        }

        // Delete related jobs records
        if err := tx.Where("employer_id = ?", existingEmployer.ID).Delete(&models.Jobs{}).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to delete jobs records",
            })
        }
    }

    // Delete related employer records
    if err := tx.Where("user_id = ?", userID).Delete(&models.Employer{}).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete employer records",
        })
    }

    var existingEmployee models.Employee
    if err := database.Db.First(&existingEmployee, "user_id = ?", userID).Error; err == nil {
        // Delete user applications
        if err := tx.Where("employee_id = ?", existingEmployee.ID).Delete(&models.Application{}).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to delete applications",
            })
        }
    }

    // Delete related employee records
    if err := tx.Where("user_id = ?", userID).Delete(&models.Employee{}).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete employee records",
        })
    }

    // Finally, delete the user
    if err := tx.Delete(&models.User{}, userID).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete user",
        })
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to commit transaction",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "User and related records deleted successfully",
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
	//fmt.Println(currRole.Role)
	currUser.Role=currRole.Role
	//fmt.Println(currUser.Role)
	database.Db.Save(&currUser)
	day:=time.Hour*24;
	claims=jtoken.MapClaims{
		"ID": currUser.ID,
		"email":currUser.Email,
		"role":currUser.Role,
		"expi":time.Now().Add(day*1).Unix(),
	}
	token2:=jtoken.NewWithClaims(jtoken.SigningMethodHS256,claims)
	t,err:=token2.SignedString([]byte(config.Secret))
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error":err.Error(),})
	}
	return c.JSON(models.LoginResponse{
		Token:t,
	})

}
