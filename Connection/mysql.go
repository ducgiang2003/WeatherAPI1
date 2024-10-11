package connection

import (
	"log"
	model "weather/Model"

	"gorm.io/driver/mysql" // Thêm dòng này để import gói mysql
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbErr error

func MySQLConnection(connectionString string) {
	Instance, dbErr = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbErr != nil {
		panic("Connection to db error")
	}
	log.Println("Connect to db success ")
}
func Migration() {
	Instance.AutoMigrate(&model.User{})
	log.Println("Migration success")
}
