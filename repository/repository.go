package repository
import(
	"errors"
	//"gorm.io/gorm"
	"github.com/DevanshS9881/Job_Portal-GO/models"
)
func Find(email,password string) (*models.User,error){
	if email=="test@gmail.com" && password=="pass1234"{
		return &models.User{
			//ID :1,
			Email: "test@gmail.com",
			Password: "pass1234",
		},nil
	}
	return nil,errors.New("user is not found")
}