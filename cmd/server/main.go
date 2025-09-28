package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pz3/internal/api"
	"pz3/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Коллекция
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	// Элемент
	mux.HandleFunc("GET /tasks/", h.GetTask)
	mux.HandleFunc("PATCH /tasks/{id}", h.UpdateDoneTask)
	mux.HandleFunc("DELETE /tasks/{id}", h.DeleteTask)

	handler := api.Cors(mux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  15 * time.Second, // Максимальное время чтения запроса
		WriteTimeout: 15 * time.Second, // Максимальное время записи ответа
		IdleTimeout:  60 * time.Second, // Максимальное время бездействия
	}

	// Канал для получения сигналов ОС
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Ждем сигнал завершения
	<-stop
	log.Println("Shutdown signal received")

	// Создаем контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Останавливаем сервер
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
