package controllers

import (
	"fmt"
	//"strings"
    "net/url"
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

func DeleteJob(c *fiber.Ctx) error {
    // Parse the job ID from the URL parameters
    jobID := database.Convert(c.Params("id"))

    // Fetch the job record from the database
    var job models.Jobs
    result := database.Db.First(&job, jobID)
    if result.Error != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid job id")
    }

    // Start a transaction
    tx := database.Db.Begin()
    if tx.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to start transaction",
        })
    }

    // Delete related applications
    if err := tx.Where("jobs_id = ?", jobID).Delete(&models.Application{}).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete related applications",
        })
    }

    // Delete the job
    if err := tx.Delete(&models.Jobs{}, jobID).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete job",
        })
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to commit transaction",
        })
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
        "data":    nil,
        "success": true,
        "message": "Successfully deleted the job and related applications",
    })
}


func GetJobByProfile(c *fiber.Ctx) error {
    profile := c.Params("profile")
	decodedProfile, err := url.QueryUnescape(profile)
    if err != nil {
        fmt.Println("Error decoding profile parameter:", err) // Debug logging
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid profile parameter",
        })
    }
    fmt.Println("Received profile:", decodedProfile) // Debug logging
   // profile = strings.TrimSpace(profile)

    var jobs []models.Jobs
    result := database.Db.Select("profile, comapny, experience, qualification, location, salary, status").
        Where("LOWER(profile) = LOWER(?)", decodedProfile).
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


 


