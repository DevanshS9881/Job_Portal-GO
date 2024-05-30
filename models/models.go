package models
import(
	"github.com/jinzhu/gorm"
)
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
	UserRole string `json:"Role"`
	City string  `json:"City"`
	Birth_Date string `json:"BirthDate"`
	Age uint `json:"Age"`
	Bio string `json:"Bio"`
	Skill string `json:"Skill"`
}
type Employer struct{
	gorm.Model
	UserID uint `json:"UserID"`
	Name string `json:"Name"`
	//User_Name string `json:"Username"`
	UserRole string `json:"Role"`
	City string  `json:"City"`
	Birth_Date string `json:"BirthDate"`
	Age uint `json:"Age"`
	Company string `json:"Company"`
	Jobs []Jobs
}
type Jobs struct{
	gorm.Model
	EmployerID uint 
	Profile string
    Comapny string
	Experience string
	//Desc string
	Location string
	Salary int64
	Status string
}
type Vacancy struct{
	gorm.Model
	JobsProfile string
	Vacancies string
}
