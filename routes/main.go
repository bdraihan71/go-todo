package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	todo := router.Group("/todo")
    {
        todo.POST("/", controllers.CreateTodo)    // Create - POST /todo/
        todo.GET("/", controllers.GetTodos)       // Read - GET /todo/
        todo.GET("/:id", controllers.GetTodo)     // Read single - GET /todo/:id
        todo.PUT("/:id", controllers.UpdateTodo)  // Update - PUT /todo/:id
        todo.DELETE("/:id", controllers.DeleteTodo) // Delete - DELETE /todo/:id
    }
	return router
}