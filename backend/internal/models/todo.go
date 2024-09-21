package models

import "time"

type Todo struct {
	ID        int64     `json:"id,omitempty"`
	Title     string    `json:"title" binding:"required"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateTodoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoRequest struct {
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
