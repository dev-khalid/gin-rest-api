package router

import (
	"github.com/dev-khalid/gin-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvents)
	server.GET("/events/:id", handlers.GetEvent)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)

}
