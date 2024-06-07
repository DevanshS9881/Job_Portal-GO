package repository

import (
	"errors"
	//"gorm.io/gorm"
	"github.com/DevanshS9881/Job_Portal-GO/database"
	"github.com/DevanshS9881/Job_Portal-GO/models"
	"github.com/DevanshS9881/Job_Portal-GO/hashPassword"
)
func Find(email,password string) (*models.User,error){
    var existingUser models.User
	existingUser.Password,_=hashpassword.HashPassword(existingUser.Password)
	// //if err:=database.Db.Where("email = ? AND password = ?", email, password).First(&existingUser).Error;err!=nil{
	// 	return nil,errors.New("user is not found")
	// }
    // //return &models.User{
	// 	//ID :1,
	// 	Email: email,
	// 	Password: password,
	// }
	if err:=database.Db.Where("email = ?", email).First(&existingUser).Error;err!=nil{
		return nil,errors.New("user is not found")
	}
	if hashpassword.VerifyPassword(password,existingUser.Password){
		return &models.User{
				//ID :1,
				Email: email,
				//Password: password,
			},nil
	}
	return nil,errors.New("user not found")
}