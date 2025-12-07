package main

import (
	"log"
	"net/http"
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
