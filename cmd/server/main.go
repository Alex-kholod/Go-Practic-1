package main

import (
	"log"
	"net/http"

	router "pz10/internal/http"
	"pz10/internal/platform/config"
)

func main() {
	cfg := config.Load()
	mux := router.Build(cfg)

	log.Println("listening on", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
