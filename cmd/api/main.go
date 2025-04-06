package main

import (
	"github.com/Liedsonfsa/API-Challenges-Points-of-Interest/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/register", routes.Register)
	server.GET("/list", routes.ListAllPoints)
	server.GET("/get-points", routes.GetPoints)

	server.Run(":3000")
}