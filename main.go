package main

import (
	"log"
	"weather/route"

	"github.com/joho/godotenv"
)

func Loadenv() {
	//Handle read file env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("There a error when load .env file : %v", err)
	}
}
func main() {
	Loadenv()
	route.Routes()
}
