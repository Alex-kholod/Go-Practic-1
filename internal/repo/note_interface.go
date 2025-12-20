package repo

import (
	"context"
	"time"

	"pz11/internal/core"
)

type NoteRepository interface {
	Create(ctx context.Context, n *core.Note) (int64, error)
	GetAll(ctx context.Context, offset, limit int) ([]*core.Note, error)
	GetAllKeyset(ctx context.Context, seenCreatedAt *time.Time, seenID *int64, limit int) ([]*core.Note, error)
	GetByID(ctx context.Context, id int64) (*core.Note, error)
	Update(ctx context.Context, n *core.Note) error
	Delete(ctx context.Context, id int64) error
	BatchGetByIDs(ctx context.Context, ids []int64) ([]*core.Note, error)
}
