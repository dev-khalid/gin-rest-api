package router

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	EventRoutes(server)
	UserRoutes(server)
}
