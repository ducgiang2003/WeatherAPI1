package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestAuth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Result": "Bang",
	})
}
