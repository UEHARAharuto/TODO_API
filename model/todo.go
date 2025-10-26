package model

import "time"

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	Title    string  `json:"title" validate:"required"` // Titleは必須
	Status   *string `json:"status"`                    // オプショナル（ポインタ型）
	Priority *int    `json:"priority"`                  // オプショナル（ポインタ型）
}

type UpdateTodoRequest struct {
	Title    *string `json:"title"`    // オプショナル（ポインタ型）
	Status   *string `json:"status"`   // オプショナル（ポインタ型）
	Priority *int    `json:"priority"` // オプショナル（ポインタ型）
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
