package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be %d", http.StatusOK)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, totalCount, len(list), "Expected cafe count to be %d, got %d", totalCount, len(list))
}

func TestMainHandlerWhenNoEmptyBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be %d", http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "Body must be not Empty")
}

func TestMainHandlerUnknownCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=unknown_city", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code to be %d", http.StatusBadRequest)

	body := responseRecorder.Body.String()
	exError := "wrong city value"
	assert.Equal(t, exError, body, "Expected response body to be '%s', got '%s'", exError, body)
}
