package service

import (
	"errors"
	"pz11/internal/core"
	"pz11/internal/repo"
	"time"
)

type NoteService struct {
	repo repo.NoteRepository
}

func NewNoteService(repo repo.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(title, content string) (*core.Note, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if len(title) > 100 {
		return nil, errors.New("title is too long (max 100 characters)")
	}

	if len(content) > 1000 {
		return nil, errors.New("content is too long (max 1000 characters)")
	}

	note := core.Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	id, err := s.repo.Create(note)
	if err != nil {
		return nil, err
	}

	note.ID = id
	return &note, nil
}

func (s *NoteService) GetNoteByID(id int64) (*core.Note, error) {
	if id <= 0 {
		return nil, errors.New("invalid note ID")
	}

	return s.repo.GetByID(id)
}

func (s *NoteService) GetAllNotes() ([]*core.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) UpdateNote(id int64, title, content string) (*core.Note, error) {
	if id <= 0 {
		return nil, errors.New("invalid note ID")
	}

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if len(title) > 100 {
		return nil, errors.New("title is too long (max 100 characters)")
	}

	if len(content) > 1000 {
		return nil, errors.New("content is too long (max 1000 characters)")
	}

	currentNote, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	updatedNote := core.Note{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: currentNote.CreatedAt,
		UpdatedAt: &now,
	}

	err = s.repo.Update(id, updatedNote)
	if err != nil {
		return nil, err
	}

	return &updatedNote, nil
}

func (s *NoteService) DeleteNote(id int64) error {
	if id <= 0 {
		return errors.New("invalid note ID")
	}

	return s.repo.Delete(id)
}
