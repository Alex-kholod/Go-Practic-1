package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"pz11/internal/core"
	"pz11/internal/core/service"
	httpx "pz11/internal/http"
	"pz11/internal/http/handlers"
	"pz11/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func newTestServer(t *testing.T, pool *pgxpool.Pool) *httptest.Server {
	t.Helper()

	noteRepo := repo.NewNoteRepoPg(pool)
	noteService := service.NewNoteService(noteRepo)
	handler := &handlers.Handler{Service: noteService}
	router := httpx.NewRouter(handler)

	return httptest.NewServer(router)
}

func createTestNote(t *testing.T, serverURL string, title, content string) *core.Note {
	t.Helper()

	reqBody := fmt.Sprintf(`{"title":"%s","content":"%s"}`, title, content)
	resp, err := http.Post(serverURL+"/api/v1/notes", "application/json", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Failed to create note: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status 201, got %d: %s", resp.StatusCode, string(body))
	}

	var note core.Note
	if err := json.NewDecoder(resp.Body).Decode(&note); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	return &note
}

func TestNoteCRUD(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	if dbHost == "" || dbUser == "" || dbName == "" {
		t.Skip("Database environment variables not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	var tableExists bool
	err = pool.QueryRow(ctx,
		"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'notes')").
		Scan(&tableExists)
	if err != nil {
		t.Fatalf("Failed to check table existence: %v", err)
	}

	if !tableExists {
		t.Fatal("Table 'notes' does not exist")
	}

	server := newTestServer(t, pool)
	defer server.Close()

	// Тест создания заметки
	t.Run("Create Note", func(t *testing.T) {
		note := createTestNote(t, server.URL, "Integration Test", "This is a test note")

		if note.ID == 0 {
			t.Error("Note ID should not be zero")
		}
		if note.Title != "Integration Test" {
			t.Errorf("Expected title 'Integration Test', got '%s'", note.Title)
		}
		if !note.CreatedAt.Before(time.Now().Add(time.Second)) {
			t.Error("CreatedAt should be in the past")
		}
	})

	// Тест получения заметки
	t.Run("Get Note", func(t *testing.T) {
		createdNote := createTestNote(t, server.URL, "Test Get", "Content for get test")

		resp, err := http.Get(fmt.Sprintf("%s/api/v1/notes/%d", server.URL, createdNote.ID))
		if err != nil {
			t.Fatalf("Failed to get note: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}

		var retrievedNote core.Note
		if err := json.NewDecoder(resp.Body).Decode(&retrievedNote); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if retrievedNote.ID != createdNote.ID {
			t.Errorf("Expected ID %d, got %d", createdNote.ID, retrievedNote.ID)
		}
		if retrievedNote.Title != createdNote.Title {
			t.Errorf("Expected title '%s', got '%s'", createdNote.Title, retrievedNote.Title)
		}
	})

	// Тест обновления заметки
	t.Run("Update Note", func(t *testing.T) {
		createdNote := createTestNote(t, server.URL, "Original Title", "Original Content")

		updateData := `{"title":"Updated Title","content":"Updated Content"}`
		req, err := http.NewRequest("PUT",
			fmt.Sprintf("%s/api/v1/notes/%d", server.URL, createdNote.ID),
			strings.NewReader(updateData))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to update note: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}

		var updatedNote core.Note
		if err := json.NewDecoder(resp.Body).Decode(&updatedNote); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if updatedNote.Title != "Updated Title" {
			t.Errorf("Expected title 'Updated Title', got '%s'", updatedNote.Title)
		}
		if updatedNote.UpdatedAt == nil {
			t.Error("UpdatedAt should be set after update")
		}
	})

	// Тест удаления заметки
	t.Run("Delete Note", func(t *testing.T) {
		createdNote := createTestNote(t, server.URL, "To Delete", "Will be deleted")

		req, err := http.NewRequest("DELETE",
			fmt.Sprintf("%s/api/v1/notes/%d", server.URL, createdNote.ID), nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to delete note: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 204, got %d: %s", resp.StatusCode, string(body))
		}

		resp, err = http.Get(fmt.Sprintf("%s/api/v1/notes/%d", server.URL, createdNote.ID))
		if err != nil {
			t.Fatalf("Failed to check deleted note: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404 for deleted note, got %d", resp.StatusCode)
		}
	})
}
