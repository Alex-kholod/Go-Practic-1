package httpx

import (
	"pz11/internal/http/handlers"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/docs/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Post("/", h.CreateNote)
			r.Get("/", h.GetAllNotes)
			r.Get("/{id}", h.GetNote)
			r.Put("/{id}", h.UpdateNote)
			r.Delete("/{id}", h.DeleteNote)
			r.Patch("/api/v1/notes/{id}", h.UpdateNote)
			r.Get("/batch", h.BatchGetNotes)
		})
	})

	return r
}
