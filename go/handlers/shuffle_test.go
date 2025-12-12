package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler := NewShuffleHandler()
	r.POST("/api/shuffle", handler.Shuffle)
	return r
}

func TestShuffle_Success(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob", "Charlie", "Dave"},
		GroupSize:    intPtr(2),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.ShuffleResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(response.Groups) < 2 {
		t.Errorf("Expected at least 2 groups, got %d", len(response.Groups))
	}
}

func TestShuffle_WithNumGroups(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"},
		NumGroups:    intPtr(3),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.ShuffleResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if len(response.Groups) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(response.Groups))
	}
}

func TestShuffle_TooFewParticipants(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob"},
		GroupSize:    intPtr(2),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "validation_error" {
		t.Errorf("Expected error 'validation_error', got '%s'", response.Error)
	}

	if response.Message != "参加者は4名以上必要です" {
		t.Errorf("Expected message '参加者は4名以上必要です', got '%s'", response.Message)
	}
}

func TestShuffle_NoGroupSizeOrNumGroups(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob", "Charlie", "Dave"},
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "グループサイズまたはグループ数を指定してください" {
		t.Errorf("Expected message 'グループサイズまたはグループ数を指定してください', got '%s'", response.Message)
	}
}

func TestShuffle_InvalidGroupSize(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob", "Charlie", "Dave"},
		GroupSize:    intPtr(0),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "グループサイズは1以上を指定してください" {
		t.Errorf("Expected message 'グループサイズは1以上を指定してください', got '%s'", response.Message)
	}
}

func TestShuffle_InvalidNumGroups(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{"Alice", "Bob", "Charlie", "Dave"},
		NumGroups:    intPtr(1),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "グループ数は2以上を指定してください" {
		t.Errorf("Expected message 'グループ数は2以上を指定してください', got '%s'", response.Message)
	}
}

func TestShuffle_InvalidJSON(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "invalid_request" {
		t.Errorf("Expected error 'invalid_request', got '%s'", response.Error)
	}
}

func TestShuffle_EmptyParticipants(t *testing.T) {
	router := setupRouter()

	reqBody := models.ShuffleRequest{
		Participants: []string{},
		GroupSize:    intPtr(2),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestShuffle_LargeGroup(t *testing.T) {
	router := setupRouter()

	participants := make([]string, 100)
	for i := 0; i < 100; i++ {
		participants[i] = string(rune('A' + i%26))
	}

	reqBody := models.ShuffleRequest{
		Participants: participants,
		GroupSize:    intPtr(5),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/shuffle", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.ShuffleResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	totalMembers := 0
	for _, group := range response.Groups {
		totalMembers += len(group)
	}

	if totalMembers != 100 {
		t.Errorf("Expected 100 total members, got %d", totalMembers)
	}
}

// Helper function to create pointer to int
func intPtr(i int) *int {
	return &i
}
