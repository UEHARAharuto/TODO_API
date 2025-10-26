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
	var req model.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// リクエストでStatusが指定されなかった場合、デフォルト値 'pending' を設定
	status := req.Status
	if status == "" {
		status = "pending"
	}

	sql := "INSERT INTO todos(title, status, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now()

	_, err := db.DB.Exec(sql, req.Title, status, now, now)
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

	sqlBase := "SELECT id, title, status, created_at, updated_at FROM todos"

	if query != "" {
		sql := sqlBase + " WHERE title LIKE ?"
		rows, err = db.DB.Query(sql, "%"+query+"%")
	} else {
		sql := sqlBase
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
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
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
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	sql := "UPDATE todos SET title = ?, status = ?, updated_at = ? WHERE id = ?"
	now := time.Now()

	result, err := db.DB.Exec(sql, req.Title, req.Status, now, id)
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
