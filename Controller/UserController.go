package controller

//Register new user
import (
	"net/http"
	connection "weather/Connection"
	model "weather/Model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Error": "Error at BindJSON" + err.Error(),
		})
		return
	}

	//HashPassword
	err = user.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error at HashPassword" + err.Error(),
		})
		return
	}

	//Create new record

	record := connection.Instance.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error when create new record" + err.Error(),
		})
		return
	}

}
