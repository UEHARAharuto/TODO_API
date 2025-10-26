package model

import "time"

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Title string `json:"title"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updatedAt"`
}
