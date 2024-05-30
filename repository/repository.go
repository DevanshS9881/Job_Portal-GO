package repository

import (
	"errors"
	//"gorm.io/gorm"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
)
func Find(email,password string) (*models.User,error){
    var existingUser models.User
	if err:=database.Db.Where("email = ? AND password = ?", email, password).First(&existingUser).Error;err!=nil{
		return nil,errors.New("user is not found")
	}
    return &models.User{
		//ID :1,
		Email: email,
		Password: password,
	},nil
}