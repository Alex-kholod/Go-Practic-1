package repo

import (
	"context"
	"errors"
	"time"

	"pz11/internal/core"
	"pz11/internal/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepoPg struct {
	pool *pgxpool.Pool
}

func NewNoteRepoPg(pool *pgxpool.Pool) *NoteRepoPg {
	return &NoteRepoPg{pool: pool}
}

func (r *NoteRepoPg) Create(ctx context.Context, n *core.Note) (int64, error) {
	var id int64
	err := r.pool.QueryRow(ctx, db.SqlInsertNote,
		n.Title, n.Content, n.CreatedAt, n.UpdatedAt,
	).Scan(&id)
	return id, err
}

func (r *NoteRepoPg) GetAll(ctx context.Context, offset, limit int) ([]*core.Note, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, content, created_at, updated_at
		 FROM notes
		 ORDER BY created_at DESC, id DESC
		 OFFSET $1 LIMIT $2
		 `, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*core.Note, 0, limit)
	for rows.Next() {
		var n core.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		nc := n
		res = append(res, &nc)
	}
	return res, rows.Err()
}

func (r *NoteRepoPg) GetByID(ctx context.Context, id int64) (*core.Note, error) {
	var n core.Note
	err := r.pool.QueryRow(ctx,
		`SELECT id, title, content, created_at, updated_at
		 FROM notes
		 WHERE id = $1`, id,
	).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *NoteRepoPg) Update(ctx context.Context, n *core.Note) error {
	ct, err := r.pool.Exec(ctx,
		`UPDATE notes
		 SET title = $1, content = $2, updated_at = $3
		 WHERE id = $4`,
		n.Title, n.Content, n.UpdatedAt, n.ID,
	)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *NoteRepoPg) Delete(ctx context.Context, id int64) error {
	ct, err := r.pool.Exec(ctx, `DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

var ErrNotFound = errors.New("note not found")

// ...

func (r *NoteRepoPg) GetAllKeyset(ctx context.Context, seenCreatedAt *time.Time, seenID *int64, limit int) ([]*core.Note, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	var (
		rows pgx.Rows
		err  error
	)

	if seenCreatedAt == nil || seenID == nil {
		rows, err = r.pool.Query(ctx, `
			SELECT id, title, content, created_at, updated_at
			FROM notes
			ORDER BY created_at DESC, id DESC
			LIMIT $1
		`, limit)
	} else {
		rows, err = r.pool.Query(ctx, `
			SELECT id, title, content, created_at, updated_at
			FROM notes
			WHERE (created_at, id) < ($1, $2)
			ORDER BY created_at DESC, id DESC
			LIMIT $3
		`, *seenCreatedAt, *seenID, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*core.Note, 0, limit)
	for rows.Next() {
		var n core.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		nc := n
		res = append(res, &nc)
	}
	return res, rows.Err()
}

func (r *NoteRepoPg) BatchGetByIDs(ctx context.Context, ids []int64) ([]*core.Note, error) {
	if len(ids) == 0 {
		return []*core.Note{}, nil
	}

	rows, err := r.pool.Query(ctx, `
		SELECT id, title, content, created_at, updated_at
		FROM notes
		WHERE id = ANY($1)
	`, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*core.Note, 0, len(ids))
	for rows.Next() {
		var n core.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		nc := n
		res = append(res, &nc)
	}
	return res, rows.Err()
}
