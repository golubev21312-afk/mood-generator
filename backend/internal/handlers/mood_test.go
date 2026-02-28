package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(h *Handler) *gin.Engine {
	r := gin.New()
	r.POST("/api/mood", h.PostMood)
	r.GET("/api/history", h.GetHistory)
	r.GET("/api/mood/:id", h.GetMoodByID)
	return r
}

// --- POST /api/mood ---

func TestPostMood_NoBody(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mood", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	assertErrorField(t, w.Body.Bytes())
}

func TestPostMood_EmptyInputField(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	body, _ := json.Marshal(map[string]string{"input": ""})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mood", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	assertErrorField(t, w.Body.Bytes())
}

func TestPostMood_MissingInputField(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	body, _ := json.Marshal(map[string]string{"wrong_field": "что-то"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mood", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	assertErrorField(t, w.Body.Bytes())
}

func TestPostMood_InvalidJSON(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mood", bytes.NewBufferString("не JSON"))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

// --- GET /api/mood/:id ---

func TestGetMoodByID_InvalidID_Letters(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/mood/abc", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
	assertErrorField(t, w.Body.Bytes())
}

func TestGetMoodByID_InvalidID_Float(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/mood/1.5", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestGetMoodByID_InvalidID_Empty(t *testing.T) {
	r := setupRouter(&Handler{DB: nil})

	// gin не смаршрутирует /api/mood/ на :id — вернёт 404 роутером
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/mood/", nil)
	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Errorf("expected non-200 for empty id, got 200")
	}
}

// --- helpers ---

// assertErrorField проверяет, что тело ответа содержит поле "error"
func assertErrorField(t *testing.T, body []byte) {
	t.Helper()
	var resp map[string]any
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("response is not JSON: %s", string(body))
	}
	if _, ok := resp["error"]; !ok {
		t.Errorf("expected 'error' field in response, got: %s", string(body))
	}
}
