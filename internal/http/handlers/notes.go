package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pz11/internal/core"
	"pz11/internal/core/service"
)

type Handler struct {
	NoteService *service.NoteService
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
	var request core.NoteCreate

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	note, err := h.NoteService.CreateNote(request.Title, request.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// GetNote godoc
// @Summary      Получить заметку
// @Tags         notes
// @Param        id   path   int  true  "ID"
// @Success      200  {object}  core.Note
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [get]
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	note, err := h.NoteService.GetNoteByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
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
	notes, err := h.NoteService.GetAllNotes()
	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
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
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var request core.NoteUpdate
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if request.Title == nil && request.Content == nil {
		http.Error(w, "At least one field (title or content) must be provided for update", http.StatusBadRequest)
		return
	}

	currentNote, err := h.NoteService.GetNoteByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	title := currentNote.Title
	content := currentNote.Content

	if request.Title != nil {
		title = *request.Title
	}
	if request.Content != nil {
		content = *request.Content
	}

	note, err := h.NoteService.UpdateNote(id, title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DeleteNote godoc
// @Summary      Удалить заметку
// @Tags         notes
// @Param        id  path  int  true  "ID"
// @Success      204  "No Content"
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [delete]
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.NoteService.DeleteNote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
