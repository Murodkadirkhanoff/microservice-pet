package main

import (
	"github.com/Murodkadirkhanoff/uiux-design-service/db"
	"github.com/Murodkadirkhanoff/uiux-design-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8070")
}
