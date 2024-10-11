// This class get user from login and generate token if it true
package controller

import (
	"net/http"
	auth1 "weather/Auth1"
	connection "weather/Connection"
	model "weather/Model"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateNewToken(context *gin.Context) {
	var request TokenRequest
	var user model.User

	//Check JSON syntax
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error at TokenController request " + err.Error(),
		})
		return
	}
	//Check Email in the DB
	record := connection.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error when find Email at TokenController(GenerateNewToken) " + err.Error(),
		})
		return
	}
	//Check password
	credentialErr := user.CheckPassword(request.Password)
	if credentialErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error at TokenController CheckPassword " + err.Error(),
		})
		return
	}
	//Generate new  token string
	tokenString, err := auth1.GenerateToken(user.Username, user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error when create new tokenString at TokenController " + err.Error(),
		})
	}
	//Print tokenString
	context.JSON(http.StatusOK, gin.H{
		"Status": "Success",
		"Token":  tokenString,
	})
}
