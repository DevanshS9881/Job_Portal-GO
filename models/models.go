package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)
type Claims struct {
	jwt.StandardClaims
}
type Roles struct{
	Role string `json:"Role"`
}
type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"Password"`
	}
type LoginResponse struct{
	Token string `json:"token"`
}
type User struct{
	gorm.Model
	Name string`json:"Name"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	Role string     `josn:"Role"`
	Employer Employer `gorm:"foreignKey:UserID"`
	Employee Employee  `gorm:"foreignKey:UserID"`
	//Jobs []Jobs
}
type Employee struct{
	gorm.Model
    UserID uint `json:"UserID"`
	Name string `json:"Name"`
	//User_Name string `json:"Username"`
	//UserRole string `json:"Role"`
	City string  `json:"City"`
	Birth_Date string `json:"BirthDate"`
	Age uint `json:"Age"`
	Bio string `json:"Bio"`
	Skill string `json:"Skill"`
	Application []Application `gorm:"foreignKey:EmployeeID"`
}
type Employer struct{
	gorm.Model
	UserID uint `json:"UserID"`
	Name string `json:"Name"`
	//User_Name string `json:"Username"`
	//UserRole string `json:"Role"`
	City string  `json:"City"`
	Birth_Date string `json:"BirthDate"`
	Age uint `json:"Age"`
	Company string `json:"Company"` 
	Jobs []Jobs 
	
}
	type Jobs struct{
		gorm.Model
		EmployerID uint `json:"Employer_ID"`
		Profile string   `json:"Profile"`
		Comapny string     `json:"Comapny"`
		Experience string   `json:"Experience"`
		Qualification string  `json:"Qualification"`
		Location string      `json:"Location"`
		Salary string        `json:"Salary"`
		Status string        `json:"Status"`
		Desc string `json:"Desc"`
		ApplicationsRecieved uint `json:"ApplicationsRecieved"`
		Application []Application `gorm:"foreignKey:JobsID"`
	}
type Vacancy struct{
	gorm.Model
	JobsProfile string
	Vacancies string
}
type Application struct{
	gorm.Model
	EmployeeID uint `json:"EmployeeID"`
	JobsID uint `json:"JobsID"`
	Letter string `json:"Letter"`
	Review string 	`json:"Review"`
}
//Jobs []Jobs //`gorm:"foreignKey:EmployerID"`