package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginRoute(t *testing.T) {
	mux := http.NewServeMux() // can also be called router

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var input map[string]string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if input["email"] == "testuser@example.com" && input["password"] == "test" {
			w.WriteHeader(http.StatusOK)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"email": "testuser@example.com", "password": "test"}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code OK, got %v", w.Code)
	}
}
