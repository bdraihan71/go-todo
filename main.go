package main

import (
	// "github.com/gin-gonic/gin"
	// "net/http"


	"go-todo/database"
	"go-todo/routes"
)

func main() {
	database.ConnectDatabase()
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "My name is Raihan no man",
	// 	})
	// })
	// router.Run(":8080")

	router := routes.SetupRouter()
	router.Run(":8080")
}
