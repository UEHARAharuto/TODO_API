package main

import (
	"TODO_API/db"
	"TODO_API/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()

	router.POST("/todos", handler.CreateTodo)
	router.GET("/todos", handler.GetTodos)
	router.PUT("/todos/:id", handler.UpdateTodo)
	router.DELETE("/todos/:id", handler.DeleteTodo)

	router.Run(":8080")
}
