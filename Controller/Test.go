package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestAuth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Result": "Bang",
	})
}
func Success(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
      <div style="
          background-color: #fff;
          padding: 40px;
          border-radius: 8px;
          box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
          text-align: center;
      ">
          <h1 style="
              color: #333;
              margin-bottom: 20px;
          ">You have Successfully signed in!</h1>
          
          </div>
      </div>
  `)))
}
