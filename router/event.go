package router

import (
	"github.com/dev-khalid/fgin-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvents)
	
} 