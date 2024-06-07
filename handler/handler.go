package handler

import (
	"context"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/DevanshS9881/Job_Portal-GO/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)
func Login(c *fiber.Ctx) error{
	loginRequest:=new(models.LoginRequest)
	if err:=c.BodyParser(loginRequest);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error(),})
	}
	user,err:=repository.Find(loginRequest.Email,loginRequest.Password)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ "error":err.Error(),})
	}
	day:=time.Hour*24;
	claims:=jtoken.MapClaims{
		"ID": user.ID,
		"email":user.Email,
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
	func Protected(c *fiber.Ctx) error{
		user:= c.Locals("user").(*jtoken.Token)
		claims:= user.Claims.(jtoken.MapClaims)
		email:= claims["email"].(string)
		return c.SendString("Welcom "+email)
	}

	func GoogleLogin(c *fiber.Ctx) error {

		url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	
		//c.Status(fiber.StatusSeeOther)
		return c.Redirect(url,fiber.StatusSeeOther)
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
	
		userData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.SendString("JSON Parsing Failed")
		}
	     var newUser models.User
		 json.Unmarshal(userData,&newUser)
		 if err:=database.Db.Where("email = ?", newUser.Email).First(&newUser).Error;err!=nil{
			result := database.Db.Create(&newUser)
			if result.Error != nil {
				c.Status(400).JSON(&fiber.Map{
					"data":    nil,
					"success": false,
					"message": result.Error,
				})
				return result.Error
			}
			if result.Error != nil {
				c.Status(400).JSON(&fiber.Map{
					"data":    nil,
					"success": false,
					"message": result.Error,
				})
				return result.Error
			}
			//return nil,errors.New("user is not found")
		}
		
	
	day:=time.Hour*24;
	claims:=jtoken.MapClaims{
		"ID": newUser.ID,
		"email":newUser.Email,
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

		 //fmt.Println(newUser)
		//return c.SendString(string(userData))
	
	}
	
