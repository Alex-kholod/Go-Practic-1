package task

import (
	"errors"
	"sync"
	"time"
	"unicode/utf8"
)

var ErrNotFound = errors.New("task not found")
var ErrLenTitle = errors.New("task title length must be >= 3 and <= 100")

type Repo struct {
	mu    sync.RWMutex
	seq   int64
	items map[int64]*Task
}

func NewRepo() *Repo {
	return &Repo{items: make(map[int64]*Task)}
}

func (r *Repo) List(page int, limit int) []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	allTasks := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		allTasks = append(allTasks, t)
	}

	total := len(allTasks)

	if total == 0 {
		return allTasks
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		return []*Task{}
	}
	if end > total {
		end = total
	}

	return allTasks[start:end]
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) (*Task, error) {
	charCount := utf8.RuneCountInString(title)
	if charCount < 3 || charCount > 100 {
		return nil, ErrLenTitle
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	return t, nil
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}
