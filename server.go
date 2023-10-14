package main

import (
	router "create_server/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetUpRouter(r)
	// r.GET("/user")
	r.Run(":5000")
}
