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
	"context"
	"log"
	"net/http"

	"pz11/internal/core/service"
	"pz11/internal/db"
	httpx "pz11/internal/http"
	"pz11/internal/http/handlers"
	"pz11/internal/repo"
)

func main() {

	ctx := context.Background()

	dbCfg, err := db.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.NewPostgresPool(ctx, dbCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repoPg := repo.NewNoteRepoPg(pool)
	svc := service.NewNoteService(repoPg)
	h := &handlers.Handler{Service: svc}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
