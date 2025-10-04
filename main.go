package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"pz4/internal/task"
	myMW "pz4/pkg/middleware"
)

func main() {
	repo, err := task.NewRepo("tasks.json")
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	h := task.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(myMW.Logger)
	r.Use(myMW.SimpleCORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(api chi.Router) {
		r.Route("/v1", func(api chi.Router) {
			api.Mount("/tasks", h.Routes())
		})
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
