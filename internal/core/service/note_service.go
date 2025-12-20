package service

import (
	"context"
	"time"

	"pz11/internal/core"
	"pz11/internal/repo"
)

type NoteService struct {
	repo repo.NoteRepository
}

func NewNoteService(r repo.NoteRepository) *NoteService {
	return &NoteService{repo: r}
}

func (s *NoteService) CreateNote(ctx context.Context, title, content string) (*core.Note, error) {
	now := time.Now().UTC()

	n := &core.Note{
		Title:     title,
		Content:   content,
		CreatedAt: now,
	}

	id, err := s.repo.Create(ctx, n)
	if err != nil {
		return nil, err
	}
	n.ID = id

	return n, nil
}

func (s *NoteService) GetNotes(ctx context.Context, offset int, limit int) ([]*core.Note, error) {
	return s.repo.GetAll(ctx, offset, limit)
}

func (s *NoteService) GetNote(ctx context.Context, id int64) (*core.Note, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *NoteService) UpdateNote(ctx context.Context, id int64, title, content *string) (*core.Note, error) {
	n, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if n == nil {
		return nil, nil
	}

	if title != nil {
		n.Title = *title
	}
	if content != nil {
		n.Content = *content
	}

	now := time.Now().UTC()
	n.UpdatedAt = &now

	if err := s.repo.Update(ctx, n); err != nil {
		return nil, err
	}

	return n, nil
}

func (s *NoteService) DeleteNote(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *NoteService) GetNotesKeyset(ctx context.Context, seenCreatedAt *time.Time, seenID *int64, limit int) ([]*core.Note, error) {
	return s.repo.GetAllKeyset(ctx, seenCreatedAt, seenID, limit)
}

func (s *NoteService) BatchGetNotes(ctx context.Context, ids []int64) ([]*core.Note, error) {
	return s.repo.BatchGetByIDs(ctx, ids)
}
