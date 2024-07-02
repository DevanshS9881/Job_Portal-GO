package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
)

func CreateJob(c *fiber.Ctx) error{
	newJob:=new(models.Jobs)
	employerID := database.Convert(c.Params("id"))
	if err:=c.BodyParser(newJob);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	var existingEmployer models.Employer
	if err:=database.Db.First(&existingEmployer,"user_id=?",employerID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "Message":"Invalid Employer ID"     ,            
		})
	}
	newJob.EmployerID=existingEmployer.ID
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

func GetJob(c *fiber.Ctx) error {
    id := database.Convert(c.Params("id"))
    var jobs []models.Jobs
    var existingEmployer models.Employer


    result1 := database.Db.First(&existingEmployer, "user_id = ?", id)
    if result1.Error != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": "Employer not found",
        })
        return result1.Error
    }

    
    result2 := database.Db.Where("employer_id = ?", existingEmployer.ID).Find(&jobs)
    if result2.Error != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": "No jobs found for this employer",
        })
        return result2.Error
    }

    return c.Status(200).JSON(&fiber.Map{
        "data":    jobs,
        "success": true,
        "message": "Successfully fetched the data",
    })
}

func AllJobs(c *fiber.Ctx) error {
    // Create a slice to hold the jobs
    var jobs []models.Jobs
    
    // Retrieve all jobs from the database
    result := database.Db.Find(&jobs)
    
    // Check for errors
    if result.Error != nil {
        return c.Status(500).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": "Failed to retrieve jobs",
        })
    }
    
    // Return the retrieved jobs
    return c.Status(200).JSON(&fiber.Map{
        "data":    jobs,
        "success": true,
        "message": "Successfully fetched all jobs",
    })
}

func DeleteJob(c *fiber.Ctx) error{
	id:=database.Convert(c.Params("id"))
	//employerID := database.Convert(c.Params("Employer_id"))
	job:=new((models.Jobs))
	//var existingEmployer models.Employer
	result:=database.Db.First(&job,id)
	if result.Error!=nil{
		c.Status(400).SendString("Invalid job id")
		return result.Error;
	}
	// var existingEmployer models.Employer
	// if err:=database.Db.First(&existingEmployer,"user_id=?",employerID).Error;err!=nil{
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"Error":err.Error,
	// 	    "Message":"Invalid Employer ID"     ,            
	// 	})
	// }
	if err:=database.Db.Where("id=?",id).Delete(&models.Jobs{}).Error;err!=nil{
		c.Status(400).JSON(&fiber.Map{
			"data":nil,
			"success":false,
			"message":"No record exists",
		})
		return err
	}
	return c.Status(200).JSON(&fiber.Map{
		"data":nil,
		"success":true,
		"message":"Successfully deleted the data",
	})	
	
}

func GetJobByProfile(c *fiber.Ctx) error{
	profile:=c.Params("profile")
	var jobs []models.Jobs
	result := database.Db.Select("profile, company, experience, qualification, location, salary, status").
		Where("profile = ?", profile).
		Find(&jobs)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(jobs)
}

func UpdateJob(c *fiber.Ctx) error{
	var job models.Jobs
	jobID := database.Convert(c.Params("id"))
	employerID := database.Convert(c.Params("Employer_id"))
	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	var existingJob models.Jobs
	if err:=database.Db.First(&existingJob,"id=?",jobID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "Message":"Invalid Job ID"     ,            
		})
	}
	var existingEmployer models.Employer
	if err:=database.Db.First(&existingEmployer,"user_id=?",employerID).Error;err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"Error":err.Error,
		    "Message":"Invalid Employer ID"     ,            
		})
	}

	result := database.Db.Model(&existingJob).Updates(job)
        if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "data":    nil,
                "success": false,
                "message": result.Error.Error(),
            })
        }
	return c.Status(200).JSON(&fiber.Map{
		"data":    existingJob,
		"success": true,
		"message": "Successfully Updated",
	})

}


 


