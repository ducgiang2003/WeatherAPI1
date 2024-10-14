package googlecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func SignInGoogle(ctx *gin.Context) {
	//Get parameter
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	//Add provider if query
	q.Add("provider", provider)
	//Update
	ctx.Request.URL.RawQuery = q.Encode()
	//Redirect Google Login Page
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}
