package router

import (
	"github.com/dev-khalid/gin-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {
	server.GET("/users", handlers.GetUsers)
	server.POST("/users", handlers.CreateUser)
	server.GET("/users/:id", handlers.GetUserById)
	server.DELETE("/users/:id", handlers.DeleteUser)
}
