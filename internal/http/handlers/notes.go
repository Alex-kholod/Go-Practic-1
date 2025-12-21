package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"pz11/internal/core/service"
	"pz11/internal/repo"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service *service.NoteService
}

// DTO
type createNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type updateNoteRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

// JSON-ошибка
type jsonError struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, jsonError{Error: msg})
}

// CreateNote godoc
// @Summary      Создать заметку
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        input  body core.NoteCreate  true  "Данные новой заметки"
// @Success      201    {object} core.Note
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /notes [post]
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	n, err := h.Service.CreateNote(r.Context(), req.Title, req.Content)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create note")
		return
	}

	writeJSON(w, http.StatusCreated, n)
}

// GetNote godoc
// @Summary      Получить заметку
// @Tags         notes
// @Param        id   path   int  true  "ID"
// @Success      200  {object}  core.Note
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [get]
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	n, err := h.Service.GetNote(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get note")
		return
	}
	if n == nil {
		writeError(w, http.StatusNotFound, "note not found")
		return
	}

	writeJSON(w, http.StatusOK, n)
}

// GetAllNotes godoc
// @Summary      Список заметок
// @Description  Возвращает список заметок
// @Tags         notes
// @Produce      json
// @Success      200    {array}  core.Note
// @Failure      500    {object} map[string]string
// @Router       /notes [get]
func (h *Handler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit := 20
	if s := q.Get("limit"); s != "" {
		v, err := strconv.Atoi(s)
		if err != nil || v <= 0 {
			writeError(w, http.StatusBadRequest, "invalid limit")
			return
		}
		if v > 100 {
			v = 100
		}
		limit = v
	}

	mode := q.Get("mode")
	if mode == "keyset" {
		// seen_created_at + seen_id могут отсутствовать (первая страница)
		var (
			seenCreatedAt *time.Time
			seenID        *int64
		)

		if s := q.Get("seen_created_at"); s != "" {
			tm, err := time.Parse(time.RFC3339, s)
			if err != nil {
				writeError(w, http.StatusBadRequest, "invalid seen_created_at (RFC3339 expected)")
				return
			}
			seenCreatedAt = &tm
		}

		if s := q.Get("seen_id"); s != "" {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil || v <= 0 {
				writeError(w, http.StatusBadRequest, "invalid seen_id")
				return
			}
			seenID = &v
		}

		// если один из курсоров задан, а второй нет — ошибка (чтобы не было нечестных кейсов)
		if (seenCreatedAt == nil) != (seenID == nil) {
			writeError(w, http.StatusBadRequest, "both seen_created_at and seen_id must be set together")
			return
		}

		notes, err := h.Service.GetNotesKeyset(r.Context(), seenCreatedAt, seenID, limit)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to list notes")
			return
		}
		writeJSON(w, http.StatusOK, notes)
		return
	}

	// default: offset
	offset := 0
	if s := q.Get("offset"); s != "" {
		v, err := strconv.Atoi(s)
		if err != nil || v < 0 {
			writeError(w, http.StatusBadRequest, "invalid offset")
			return
		}
		offset = v
	}

	notes, err := h.Service.GetNotes(r.Context(), offset, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list notes")
		return
	}
	writeJSON(w, http.StatusOK, notes)
}

// UpdateNote godoc
// @Summary      Обновить заметку
// @Description  Обновляет заметку
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id     path   int        true  "ID заметки"
// @Param        input  body   core.NoteUpdate true  "Поля для обновления"
// @Success      200    {object}  core.Note
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /notes/{id} [put]
func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	var req updateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	n, err := h.Service.UpdateNote(r.Context(), id, req.Title, req.Content)
	if err != nil {
		if err == repo.ErrNotFound {
			writeError(w, http.StatusNotFound, "note not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to update note")
		return
	}
	if n == nil {
		writeError(w, http.StatusNotFound, "note not found")
		return
	}

	writeJSON(w, http.StatusOK, n)
}

// DeleteNote godoc
// @Summary      Удалить заметку
// @Tags         notes
// @Param        id  path  int  true  "ID"
// @Success      204  "No Content"
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [delete]
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDParam(w, r)
	if !ok {
		return
	}

	if err := h.Service.DeleteNote(r.Context(), id); err != nil {
		if err == repo.ErrNotFound {
			writeError(w, http.StatusNotFound, "note not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete note")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseIDParam(w http.ResponseWriter, r *http.Request) (int64, bool) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return 0, false
	}

	return id, true
}

func (h *Handler) BatchGetNotes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	raw := q["ids"]
	if len(raw) == 0 {
		if s := q.Get("ids"); s != "" {
			raw = []string{s}
		}
	}
	if len(raw) == 0 {
		writeError(w, http.StatusBadRequest, "ids is required")
		return
	}

	parts := make([]string, 0, 128)
	for _, chunk := range raw {
		for _, p := range strings.Split(chunk, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				parts = append(parts, p)
			}
		}
	}

	if len(parts) == 0 {
		writeError(w, http.StatusBadRequest, "ids is required")
		return
	}
	if len(parts) > 200 { // защита + лимит длины URL
		writeError(w, http.StatusBadRequest, "too many ids (max 200)")
		return
	}

	ids := make([]int64, 0, len(parts))
	seen := make(map[int64]struct{}, len(parts))
	for _, s := range parts {
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil || id <= 0 {
			writeError(w, http.StatusBadRequest, "invalid id in ids")
			return
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}

	notes, err := h.Service.BatchGetNotes(r.Context(), ids)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to batch get notes")
		return
	}
	writeJSON(w, http.StatusOK, notes)
}
