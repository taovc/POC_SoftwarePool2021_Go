package main

import (
	"SoftwareGoDay2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Routers(r)
	r.Run("localhost:8080")
}
