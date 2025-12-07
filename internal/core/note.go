package core

import "time"

type Note struct {
	ID        int64      `json:"id" example:"1"`
	Title     string     `json:"title" example:"Моя заметка"`
	Content   string     `json:"content" example:"Текст заметки"`
	CreatedAt time.Time  `json:"createdAt" example:"2024-01-15T10:30:00Z"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" example:"2024-01-15T14:45:00Z"`
}

type NoteCreate struct {
	Title   string `json:"title" validate:"required,min=1,max=100" example:"Новая заметка"`
	Content string `json:"content" validate:"max=1000" example:"Текст заметки"`
}

type NoteUpdate struct {
	Title   *string `json:"title,omitempty" validate:"omitempty,min=1,max=100" example:"Обновлено"`
	Content *string `json:"content,omitempty" validate:"omitempty,max=1000" example:"Новый текст"`
}
