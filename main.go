package main

import (
	"github.com/dev-khalid/gin-rest-api/router"
	"github.com/gin-gonic/gin"
)
func main() {
	server := gin.Default()
	router.Routes(server)

	server.Run(":8080")
}
