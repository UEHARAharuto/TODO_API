package handler

import (
	"TODO_API/db"
	"TODO_API/model"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	sql := "INSERT INTO todos(title, created_at, updated_at) VALUES (?, ?, ?)"
	now := time.Now()

	_, err := db.DB.Exec(sql, todo.Title, now, now)
	if err != nil {
		log.Printf("ERROR: Failed to create todo: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func GetTodos(c *gin.Context) {
	query := c.Query("title")

	var rows *sql.Rows
	var err error

	if query != "" {
		sql := "SELECT id, title, created_at, updated_at FROM todos WHERE title LIKE ?"
		rows, err = db.DB.Query(sql, "%"+query+"%")
	} else {
		sql := "SELECT id, title, created_at, updated_at FROM todos"
		rows, err = db.DB.Query(sql)
	}

	if err != nil {
		log.Printf("ERROR: Failed to get todos: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Printf("ERROR: Failed to scan todo row: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.Status(http.StatusNotImplemented)
		return
	}

	sql := "UPDATE todos SET title = ?, updated_at = ? WHERE id = ?"
	now := time.Now()

	result, err := db.DB.Exec(sql, todo.Title, now, id)
	if err != nil {
		log.Printf("ERROR: Failed to update todo: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("ERROR: Failed to get rows affected on update: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}
	if rowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	sql := "DELETE FROM todos WHERE id = ?"

	result, err := db.DB.Exec(sql, id)
	if err != nil {
		log.Printf("ERROR: Failed to delete todo: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("ERROR: Failed to get rows affected on delete: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}
	if rowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}
