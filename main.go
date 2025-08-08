package main

import (
	_ "github.com/dev-khalid/gin-rest-api/config"
	"github.com/dev-khalid/gin-rest-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	router.Routes(server)

	server.Run(":8080")
}
