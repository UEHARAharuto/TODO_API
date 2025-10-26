package handler

import (
	"TODO_API/db"
	"TODO_API/model"
	"database/sql"
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

	priority := todo.Priority
	if priority == 0 {
		priority = 100
	}

	sql := "INSERT INTO todos(title, priority, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now()

	_, err := db.DB.Exec(sql, todo.Title, priority, now, now)
	if err != nil {
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
		sql := sqlBase + " WHERE title LIKE ? " + sqlOrder
		rows, err = db.DB.Query(sql, "%"+query+"%")
	} else {
		sql := sqlBase + " " + sqlOrder
		rows, err = db.DB.Query(sql)
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
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

	priority := todo.Priority
	if priority == 0 {
		priority = 100
	}

	sql := "UPDATE todos SET title = ?, priority = ?, updated_at = ? WHERE id = ?"
	now := time.Now()

	result, err := db.DB.Exec(sql, todo.Title, priority, now, id)
	if err != nil {
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
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
		c.Status(http.StatusNotImplemented)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.Status(http.StatusNotImplemented)
		return
	}
	if rowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}
