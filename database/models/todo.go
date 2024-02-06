package models

import (
    "gorm.io/gorm"
)

type Todo struct {
    gorm.Model
    Title           string `form:"title" json:"title"  binding:"required"`
    Description     string `form:"description" json:"description" binding:"required"`
	IsComplete  bool   `form:"is_complete" json:"is_complete" gorm:"default:false"`
}