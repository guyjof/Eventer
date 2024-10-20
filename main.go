package main

import (
	"example.com/eventer-api/db"
	"example.com/eventer-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterHealthRoutes(server)
	routes.RegisterApplicationRoutes(server)

	server.Run(":8080")
}
