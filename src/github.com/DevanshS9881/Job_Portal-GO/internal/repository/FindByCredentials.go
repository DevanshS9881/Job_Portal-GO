pacakge repository
import(
	"errors"
	"github.com/DevanshS9881/Job_Portal-GO/internal/models"
)
func find(email,password string) (*models.User,error){
	if email=="test@gmail.com" && password="pass1234"{
		return &models.User{
			ID:1,
			Email:test@gmail.com,
			Password:pass1234,
		},nil
	}
	return nil,errors.New("User is Not Found")
}