package repo

import (
	"errors"
	"pz11/internal/core"
	"sync"
)

type NoteRepository interface {
	Create(note core.Note) (int64, error)
	GetByID(id int64) (*core.Note, error)
	GetAll() ([]*core.Note, error)
	Update(id int64, note core.Note) error
	Delete(id int64) error
}

type NoteRepoMem struct {
	mu    sync.RWMutex
	notes map[int64]*core.Note
	next  int64
}

func NewNoteRepoMem() *NoteRepoMem {
	return &NoteRepoMem{notes: make(map[int64]*core.Note), next: 0}
}

func (r *NoteRepoMem) Create(n core.Note) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.next++
	n.ID = r.next
	noteCopy := n
	r.notes[n.ID] = &noteCopy
	return n.ID, nil
}

func (r *NoteRepoMem) GetByID(id int64) (*core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, exists := r.notes[id]
	if !exists {
		return nil, errors.New("note not found")
	}

	noteCopy := *note
	return &noteCopy, nil
}

func (r *NoteRepoMem) GetAll() ([]*core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	notes := make([]*core.Note, 0, len(r.notes))
	for _, note := range r.notes {
		noteCopy := *note
		notes = append(notes, &noteCopy)
	}
	return notes, nil
}

func (r *NoteRepoMem) Update(id int64, note core.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.notes[id]; !exists {
		return errors.New("note not found")
	}

	note.CreatedAt = r.notes[id].CreatedAt
	r.notes[id] = &note
	return nil
}

func (r *NoteRepoMem) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.notes[id]; !exists {
		return errors.New("note not found")
	}

	delete(r.notes, id)
	return nil
}
