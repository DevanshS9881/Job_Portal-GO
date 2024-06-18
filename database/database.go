package database

import (
	"log"
	"strconv"

	"github.com/DevanshS9881/Job_Portal-GO/config"
	"github.com/DevanshS9881/Job_Portal-GO/models"

	//"github.com/jinzhu/gorm"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
    func Convert(port string) uint{
    u64, err := strconv.ParseUint(port, 10, 64)
    if err != nil {
        log.Fatal("Error:", err)
    }
    return uint(u64)
}
var Db *gorm.DB

var dsn string =fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",config.Load("host"),Convert(config.Load("port")), config.Load("user"), config.Load("password"), config.Load("dbname"))

func InitDB() error{
	
	var err error
		Db,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		return err
	}
	fmt.Println("Successfully connected to the database")
	Db.AutoMigrate(&models.User{},&models.Jobs{},&models.Employer{},&models.Employee{},&models.Application{})
    return nil
}	

