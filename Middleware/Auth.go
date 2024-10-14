package middleware

import (
	"net/http"
	auth1 "weather/Auth1"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the token from the header
		// If the token is empty, return an error
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Token is empty at auth in Middleware" + err.Error(),
			})
			ctx.Abort()
		}

		//else token is no empty check valid
		err = auth1.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Token is invalid at auth in Middleware" + err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
