package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pz3/internal/storage"
	"strconv"
	"testing"
)

// Вспомогательная функция для установки ID в пути
func withTaskID(r *http.Request, id int64) *http.Request {
	r.SetPathValue("id", strconv.FormatInt(id, 10))
	return r
}

func TestTaskCRUD(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)
	var taskID int64

	t.Run("1. CREATE - создаем задачу", func(t *testing.T) {
		jsonData := `{"title":"Test Task"}`
		req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(jsonData))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		h.CreateTask(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("CREATE: статус код = %v, ожидался %v", rr.Code, http.StatusCreated)
		}

		var task map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &task)

		taskID = int64(task["id"].(float64))

		if task["title"] != "Test Task" {
			t.Errorf("CREATE: название задачи = %v, ожидалось %v", task["title"], "Test Task")
		}
		if task["done"] != false {
			t.Errorf("CREATE: статус выполнения = %v, ожидался %v", task["done"], false)
		}
	})

	t.Run("2. READ - получаем созданную задачу", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks/"+strconv.FormatInt(taskID, 10), nil)
		req = withTaskID(req, taskID)
		rr := httptest.NewRecorder()

		h.GetTask(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("READ: статус код = %v, ожидался %v", rr.Code, http.StatusOK)
		}

		var task map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &task)

		if task["id"].(float64) != float64(taskID) {
			t.Errorf("READ: ID задачи = %v, ожидался %v", task["id"], taskID)
		}
		if task["title"] != "Test Task" {
			t.Errorf("READ: название задачи = %v, ожидалось %v", task["title"], "Test Task")
		}
	})

	t.Run("3. UPDATE - отмечаем задачу как выполненную", func(t *testing.T) {
		req := httptest.NewRequest("PATCH", "/tasks/"+strconv.FormatInt(taskID, 10), nil)
		req = withTaskID(req, taskID)
		rr := httptest.NewRecorder()

		h.UpdateDoneTask(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("UPDATE: статус код = %v, ожидался %v", rr.Code, http.StatusOK)
		}

		var task map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &task)

		if task["done"] != true {
			t.Errorf("UPDATE: статус выполнения = %v, ожидался %v", task["done"], true)
		}
	})

	t.Run("4. DELETE - удаляем задачу", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/tasks/"+strconv.FormatInt(taskID, 10), nil)
		req = withTaskID(req, taskID)
		rr := httptest.NewRecorder()

		h.DeleteTask(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("DELETE: статус код = %v, ожидался %v", rr.Code, http.StatusNoContent)
		}
	})

	t.Run("5. VERIFY - проверяем что задача удалена", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks/"+strconv.FormatInt(taskID, 10), nil)
		req = withTaskID(req, taskID)
		rr := httptest.NewRecorder()

		h.GetTask(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("VERIFY: статус код = %v, ожидался %v", rr.Code, http.StatusNotFound)
		}
	})
}
