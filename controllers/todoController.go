package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo/database/models"
	"net/http"
	"go-todo/database"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	
	if err:= c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully created", "data": todo})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	if err := database.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})

	// var todos []models.Todo
    // isComplete := c.DefaultQuery("is_complete", "")

    // db := database.DB
    // if isComplete != "" {
    //     var complete bool
    //     if isComplete == "true" {
    //         complete = true
    //     } else if isComplete == "false" {
    //         complete = false
    //     } else {
    //         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid is_complete value"})
    //         return
    //     }
    //     db = db.Where("is_complete = ?", complete)
    // }

    // if err := db.Find(&todos).Error; err != nil {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
    //     return
    // }

    // c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully updated", "data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.Delete(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully deleted"})
}