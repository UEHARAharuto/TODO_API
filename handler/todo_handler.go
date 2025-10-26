package handler

import (
	"TODO_API/db"
	"TODO_API/model"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateTodo(c *gin.Context) {
	var req model.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := "INSERT INTO todos(title, status, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now()

	status := req.Status
	if status == "" {
		status = "pending"
	}

	_, err := db.DB.Exec(sql, req.Title, status, now, now)
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	priority := todo.Priority
	if priority == 0 {
		priority = 100
	}

	sql := "INSERT INTO todos(title, priority, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now()

	_, err := db.DB.Exec(sql, todo.Title, priority, now, now)
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

	sqlBase := "SELECT id, title, priority, created_at, updated_at FROM todos"
	sqlOrder := "ORDER BY priority ASC"

	if query != "" {
		sql := "SELECT id, title, status, created_at, updated_at FROM todos WHERE title LIKE ?"
		rows, err = db.DB.Query(sql, "%"+query+"%")
	} else {
		sql := "SELECT id, title, status, created_at, updated_at FROM todos"
		sql := sqlBase + " WHERE title LIKE ? " + sqlOrder
		rows, err = db.DB.Query(sql, "%"+query+"%")
	} else {
		sql := sqlBase + " " + sqlOrder
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
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Printf("ERROR: Failed to scan todo row: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	var responses []model.TodoResponse
	for _, todo := range todos {
		responses = append(responses, model.TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Status:    todo.Status,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"todos": responses})
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := "UPDATE todos SET title = ?, status = ?, updated_at = ? WHERE id = ?"
	now := time.Now()

	result, err := db.DB.Exec(sql, req.Title, req.Status, now, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.Status(http.StatusNotImplemented)
		return
	}

	priority := todo.Priority
	if priority == 0 {
		priority = 100
	}

	sql := "UPDATE todos SET title = ?, priority = ?, updated_at = ? WHERE id = ?"
	now := time.Now()

	result, err := db.DB.Exec(sql, todo.Title, priority, now, id)
	if err != nil {
		log.Printf("ERROR: Failed to update todo: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.Status(http.StatusInternalServerError)
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
		c.Status(http.StatusInternalServerError)
		log.Printf("ERROR: Failed to delete todo: %v", err)
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.Status(http.StatusInternalServerError)
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
