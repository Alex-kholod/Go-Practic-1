package core

import "time"

type Note struct {
	ID        int64      `json:"id" example:"1"`
	Title     string     `json:"title" example:"Моя заметка"`
	Content   string     `json:"content" example:"Текст заметки"`
	CreatedAt time.Time  `json:"createdAt" example:"2024-01-15T10:30:00Z"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" example:"2024-01-15T14:45:00Z"`
}
