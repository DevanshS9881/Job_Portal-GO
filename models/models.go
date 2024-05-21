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
	ID uint
	Email string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
	//Role string
	//Profile Profile
	//Jobs []Jobs
}
type Profile struct{
    gorm.Model
	Name string
	User_Name string
	Location string
	UserID uint
}
type Jobs struct{
	gorm.Model
	Post string
    Comapny string
	Profile string
	Place string
	Salary int64
	Status string
	UserId uint
}