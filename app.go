package project

import (
	"github.com/gin-gonic/gin"
)

func App(port string) {
	router := gin.Default()

	tests := router.Group("/api/test")
	{
		tests.GET("/", controllers.GetTests)
		tests.POST("/create", controllers.CreateTest)

	}

	router.Run(port)

}
