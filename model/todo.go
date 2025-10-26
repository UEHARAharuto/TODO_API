package model

import "time"

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type UpdateTodoRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
