// Package main Notes API server.
//
// @title           Notes API
// @version         1.0
// @description     REST API для заметок (CRUD).
// @contact.name    Alexey
// @contact.email   kholodkov.a.d@edu.mirea.ru
// @host            localhost:8080
// @BasePath        /api/v1
package main

import (
	"log"
	"net/http"
	_ "pz11/docs" // Важно: добавьте этот импорт
	"pz11/internal/core/service"
	httpx "pz11/internal/http"
	"pz11/internal/http/handlers"
	"pz11/internal/repo"
)

func main() {
	noteRepo := repo.NewNoteRepoMem()
	noteService := service.NewNoteService(noteRepo)

	h := &handlers.Handler{NoteService: noteService}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
