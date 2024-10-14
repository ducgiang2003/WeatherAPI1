package controller

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowGoogleHome(c *gin.Context) {
	index, err := template.ParseFiles("View/SignInGoogle.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error at showGoogleHome " + err.Error(),
		})
		c.Abort()
		return
	}
	err = index.Execute(c.Writer, gin.H{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}
func Result(c *gin.Context) {
	index, err := template.ParseFiles("View/Success.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error at Success " + err.Error(),
		})
		c.Abort()
		return
	}
	err = index.Execute(c.Writer, gin.H{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}
