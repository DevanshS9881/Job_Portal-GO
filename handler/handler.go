package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	//"fmt"
	//"io/ioutil"
	"net/http"
	"time"

	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/DevanshS9881/Job_Portal-GO/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := repository.Find(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	day := time.Hour * 24
	fmt.Println(user.ID)
	fmt.Println(user.Role)
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"expi":  time.Now().Add(day * 1).Unix(),
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}
func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome " + email)
}

	func GoogleLogin(c *fiber.Ctx) error {

		url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

		//c.Status(fiber.StatusSeeOther)
		return c.Redirect(url, fiber.StatusSeeOther)
		//	return c.JSON(url)	
	}

	func GoogleCallback(c *fiber.Ctx) error {
		state := c.Query("state")
		if state != "randomstate" {
			return c.SendString("States don't Match!!")
		}

		code := c.Query("code")

		googlecon := config.GoogleConfig()

		token, err := googlecon.Exchange(context.Background(), code)
		if err != nil {
			return c.SendString("Code-Token Exchange Failed")
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return c.SendString("Cannot retrieve the user data")
		}
		defer resp.Body.Close()

		userData, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.SendString("JSON Parsing Failed")
		}
		var newUser models.User
		json.Unmarshal(userData, &newUser)
		link:="https://job-portal-go.onrender.com/homepage.html"
		if err := database.Db.Where("email = ?", newUser.Email).First(&newUser).Error; err != nil {
			result := database.Db.Create(&newUser)
			link="https://job-portal-go.onrender.com/protected.html"
			if result.Error != nil {
				c.Status(400).JSON(&fiber.Map{
					"data":    nil,
					"success": false,
					"message": result.Error,
				})
				return result.Error
			}
			newUser.Role = "employee"
			database.Db.Save(&newUser)
		}


		day := time.Hour * 24
		claims := jtoken.MapClaims{
			"ID":    newUser.ID,
			"email": newUser.Email,
			"role":  newUser.Role,
			"expi":  time.Now().Add(day * 1).Unix(),
		}
		token2 := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
		t, err := token2.SignedString([]byte(config.Secret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    t,
			Expires:  time.Now().Add(day * 1),
			Domain: ".onrender.com",
			Path:     "/",
			Secure:   true,
			SameSite: "None",
		})
	
		// Redirect to the protected page
		return c.Redirect(link, fiber.StatusSeeOther)
		// fmt.Println()
		// return c.JSON(models.LoginResponse{
		// 	Token: t,
		// })

		//fmt.Println(newUser)
		//return c.SendString(string(userData))

	}
