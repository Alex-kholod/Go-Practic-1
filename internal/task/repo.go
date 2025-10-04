package task

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
	"unicode/utf8"
)

var ErrNotFound = errors.New("task not found")
var ErrLenTitle = errors.New("task title length must be >= 3 and <= 100")

type Repo struct {
	mu       sync.RWMutex
	seq      int64
	items    map[int64]*Task
	filePath string
}

func NewRepo(filePath string) (*Repo, error) {
	repo := &Repo{
		items:    make(map[int64]*Task),
		filePath: filePath,
		seq:      0,
	}

	if err := repo.loadFromFile(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *Repo) loadFromFile() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return r.createFileLocked()
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		r.items = make(map[int64]*Task)
		r.seq = 0
		return nil
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	r.items = make(map[int64]*Task)
	r.seq = 0
	for _, task := range tasks {
		r.items[task.ID] = task
		if task.ID > r.seq {
			r.seq = task.ID
		}
	}

	return nil
}

func (r *Repo) createFileLocked() error {
	dirPath := filepath.Dir(r.filePath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	emptyData := []byte("[]")
	return os.WriteFile(r.filePath, emptyData, 0644)
}

func (r *Repo) saveToFileLocked() error {
	tasks := make([]*Task, 0, len(r.items))
	for _, task := range r.items {
		tasks = append(tasks, task)
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *Repo) List(page int, limit int, doneFilter bool) []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	allTasks := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		if doneFilter {
			if t.Done {
				allTasks = append(allTasks, t)
			}
		} else {
			allTasks = append(allTasks, t)
		}
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

	if err := r.saveToFileLocked(); err != nil {
		delete(r.items, t.ID)
		r.seq--
		return nil, err
	}

	return t, nil
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	charCount := utf8.RuneCountInString(title)
	if charCount < 3 || charCount > 100 {
		return nil, ErrLenTitle
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}

	oldTitle := t.Title
	oldDone := t.Done
	oldUpdatedAt := t.UpdatedAt

	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()

	if err := r.saveToFileLocked(); err != nil {
		t.Title = oldTitle
		t.Done = oldDone
		t.UpdatedAt = oldUpdatedAt
		return nil, err
	}

	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.items[id]
	if !ok {
		return ErrNotFound
	}

	deletedTask := t
	delete(r.items, id)

	if err := r.saveToFileLocked(); err != nil {
		r.items[id] = deletedTask
		return err
	}

	return nil
}
