package main

import (
	"github.com/gin-gonic/gin"
	"log"
	connection "weather/Connection"
	route "weather/Route"

	"github.com/joho/godotenv"
)

var mysqlConString string = "root:@tcp(localhost:3306)/weather?parseTime=true"

func Loadenv() {
	//Handle read file env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("There a error when load .env file : %v", err)
	}
}
func main() {
	Loadenv()
	connection.MySQLConnection(mysqlConString)
	connection.Migration()

	route.Routes()
	router := gin.Default()
	router.Run(":8000")
}
