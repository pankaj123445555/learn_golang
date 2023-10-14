package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Resource struct {
	ID   int
	Name string
}

var resources []Resource

func SetUpRouter(r *gin.Engine) {

	r.GET("/api/user", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "API works successfully",
		})
	})

	r.POST("/api/send-data", func(c *gin.Context) {

		fmt.Println("finally it is called kuch bhi")
		var resource Resource
		if err := c.ShouldBind(&resource); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		resources = append(resources, resource)
		c.JSON(200, gin.H{
			"resource": resource,
		})
	})

}
