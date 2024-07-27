package controllers

import (
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/gofiber/fiber/v2"
     "github.com/unidoc/unipdf/v3/model"
       "github.com/unidoc/unipdf/v3/extractor"
       "os"
)

func Apply(c *fiber.Ctx) error {
    newAppl := new(models.Application)
    employeeID := database.Convert(c.Params("Emid"))
    jobID := database.Convert(c.Params("jobID"))
    if err := c.BodyParser(newAppl); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "Error": err.Error(),
        })
    }

    var existingEmployee models.Employee
    if err := database.Db.First(&existingEmployee, "user_id = ?", employeeID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Invalid Employee",
        })
    }

    var existingJob models.Jobs
    if err := database.Db.First(&existingJob, "id = ?", jobID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Invalid Job",
        })
    }

    newAppl.EmployeeID = existingEmployee.ID
    newAppl.JobsID = jobID
    result := database.Db.Save(&newAppl)

    if result.Error != nil {
        return c.Status(400).JSON(fiber.Map{
            "data":    nil,
            "success": false,
            "message": result.Error.Error(), // Convert error to string explicitly
        })
    }

    existingJob.ApplicationsRecieved++
    if err := database.Db.Save(&existingJob).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Failed to update job applications count",
        })
    }

    return c.Status(200).JSON(fiber.Map{
        "data":    newAppl,
        "success": true,
        "message": "Application submitted successfully",
    })
}


func GetApplicationsByEmployee(c *fiber.Ctx) error {
	employeeID := database.Convert(c.Params("id"))

	var employee models.Employee
	if err := database.Db.First(&employee, "user_id = ?", employeeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Employee not found",
		})
	}

	var applications []models.Application
	if err := database.Db.Preload("Jobs").Where("employee_id = ?", employee.ID).Find(&applications).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Could not retrieve applications",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"applications": applications,
	})
}

func uploadHandler(c *fiber.Ctx) error {

    newAppl := new(models.Application)
    employeeID := database.Convert(c.Params("Emid"))
    jobID := database.Convert(c.Params("jobID"))
	// Handle file upload
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to get file")
	}

	// Save the file
	filePath := "./uploads/" + file.Filename
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
	}

	// Process PDF
	text, err := extractTextFromPDF(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to process PDF")
	}

    var existingEmployee models.Employee
    if err := database.Db.First(&existingEmployee, "user_id = ?", employeeID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Invalid Employee",
        })
    }

    var existingJob models.Jobs
    if err := database.Db.First(&existingJob, "id = ?", jobID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Invalid Job",
        })
    }

    newAppl.EmployeeID = existingEmployee.ID
    newAppl.JobsID = jobID
    newAppl.Letter=text
    result := database.Db.Save(&newAppl)

    if result.Error != nil {
        return c.Status(400).JSON(fiber.Map{
            "data":    nil,
            "success": false,
            "message": result.Error.Error(), // Convert error to string explicitly
        })
    }

    existingJob.ApplicationsRecieved++
    if err := database.Db.Save(&existingJob).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{
            "Error":   err.Error(),
            "Message": "Failed to update job applications count",
        })
    }

    return c.Status(200).JSON(fiber.Map{
        "data":    newAppl,
        "success": true,
        "message": "Application submitted successfully",
    })
}


	// Return extracted text
// 	return c.SendString(text)
// }

func extractTextFromPDF(filePath string) (string, error) {

    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

	// Ensure you are using the correct initialization method
	pdfReader,err := model.NewPdfReader(file)
	if err != nil {
		return "", err
	}

	var text string
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	for pageIndex := 1; pageIndex <= numPages; pageIndex++ {
		page, err := pdfReader.GetPage(pageIndex)
		if err != nil {
			return "", err
		}

		textExtractor,err := extractor.New(page)
        if err!=nil{
            return "",err
        }
        pageText,err:=textExtractor.ExtractText()
        if err!=nil{
            return "",err
        }

 
		text += pageText
	}

	return text, nil
}