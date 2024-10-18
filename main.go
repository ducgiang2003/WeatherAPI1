package main

import (
	"log"
	connection "weather/Connection"
	"weather/Route"

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

	Route.Routes()

}
