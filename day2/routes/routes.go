package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func world1(c *gin.Context) {
	c.String(http.StatusOK, "%s", "hello")
}

func Routers(r *gin.Engine) {
	r.GET("/repeat-my-query", world1)
	r.GET("/repeat-my-param/:message", world1)
	r.POST("/repeat-my-body", world1)
	r.GET("/repeat-my-header", world1)
	r.GET("/repeat-my-cookie", world1)
}
