package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", handler.Ping)
		api.GET("/users", handler.GetUsers)
		api.POST("/users", handler.CreateUser)
	}

	return r
}