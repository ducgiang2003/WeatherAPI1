package googlecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func CallBackGoogle(ctx *gin.Context) {
	//Get param and query
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()

	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	_, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Eroor when callback in GoogleController " + err.Error(),
		})
		return
	}
	//Redirect to success login
	ctx.Redirect(http.StatusTemporaryRedirect, "/success")

}
