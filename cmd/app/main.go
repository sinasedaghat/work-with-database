package main

import (
	"example.url/DB/internal/database"
	"example.url/DB/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.PQConnection(".env")
	database.PQInitialize()

	server := gin.Default()

	// routes.RegisterRoutes(server)
	router.RegisterRoutes(server)

	server.Run(":8080")
}
