package configs

import (
	"log"

	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:shahil@tcp(127.0.0.1:3306)/GOLANG_UNIT_TESTS?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic("Error connecting to the Database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("Error auto migrating schemas", err)
	}

	DB = db
}
