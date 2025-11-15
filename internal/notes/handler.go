package notes

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Handler struct{ repo *Repo }

func NewHandler(r *Repo) *Handler { return &Handler{repo: r} }

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.create)
	r.Get("/", h.list)
	r.Get("/stats", h.stats)
	r.Get("/search", h.search)
	r.Get("/{id}", h.get)
	r.Patch("/{id}", h.patch)
	r.Delete("/{id}", h.del)
	return r
}

func reqCtx(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), 5*time.Second)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Title     string     `json:"title"`
		Content   string     `json:"content"`
		ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid_json"})
		return
	}

	if err := json.Unmarshal(body, &in); err != nil || in.Title == "" {
		writeJSON(w, 400, map[string]string{"error": "invalid_json_or_title"})
		return
	}

	c, cancel := reqCtx(r)
	defer cancel()
	n, err := h.repo.Create(c, in.Title, in.Content, in.ExpiresAt)
	if err != nil {
		writeJSON(w, 409, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, n)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c, cancel := reqCtx(r)
	defer cancel()
	n, err := h.repo.ByID(c, id)
	if errors.Is(err, ErrNotFound) {
		writeJSON(w, 404, map[string]string{"error": "not_found"})
		return
	}
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, n)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	after := r.URL.Query().Get("after")

	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if limit <= 0 || limit > 200 {
		limit = 20
	}

	c, cancel := reqCtx(r)
	defer cancel()
	items, err := h.repo.List(c, q, limit, after)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, items)
}

func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		writeJSON(w, 400, map[string]string{"error": "query_required"})
		return
	}

	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	c, cancel := reqCtx(r)
	defer cancel()
	items, err := h.repo.Search(c, query, limit)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, items)
}

func (h *Handler) stats(w http.ResponseWriter, r *http.Request) {
	c, cancel := reqCtx(r)
	defer cancel()

	stats, err := h.repo.GetStats(c)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, 200, stats)
}

func (h *Handler) patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in struct {
		Title     *string    `json:"title"`
		Content   *string    `json:"content"`
		ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid_json"})
		return
	}

	if len(body) == 0 {
		writeJSON(w, 400, map[string]string{"error": "empty_body"})
		return
	}

	if err := json.Unmarshal(body, &in); err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid_json"})
		return
	}

	if in.Title == nil && in.Content == nil && in.ExpiresAt == nil {
		writeJSON(w, 400, map[string]string{"error": "no_fields_to_update"})
		return
	}

	c, cancel := reqCtx(r)
	defer cancel()
	n, err := h.repo.Update(c, id, in.Title, in.Content, in.ExpiresAt)
	if errors.Is(err, ErrNotFound) {
		writeJSON(w, 404, map[string]string{"error": "not_found"})
		return
	}
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, n)
}

func (h *Handler) del(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c, cancel := reqCtx(r)
	defer cancel()
	if err := h.repo.Delete(c, id); errors.Is(err, ErrNotFound) {
		writeJSON(w, 404, map[string]string{"error": "not_found"})
		return
	} else if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(204)
}
