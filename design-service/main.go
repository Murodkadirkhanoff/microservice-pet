package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// db.Init()
	server := gin.Default()
	server.Run(":8070")
}