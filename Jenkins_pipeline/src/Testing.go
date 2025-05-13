package main

import (
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetBookById(t *testing.T) {
    // Setup the Gin router
    r := gin.Default()
    r.GET("/books/:id", getBookById)

    // Test: Book with ID "1"
    req, _ := http.NewRequest("GET", "/books/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the status code
    assert.Equal(t, http.StatusOK, w.Code)

    // Check the response body
    expected := `{"id":"1","title":"Book One","author":"Author One"}`
    assert.JSONEq(t, expected, w.Body.String())

    // Test: Book with ID "3" (Non-existing book)
    req, _ = http.NewRequest("GET", "/books/3", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the status code
    assert.Equal(t, http.StatusNotFound, w.Code)

    // Check the response body
    expectedNotFound := `{"message":"Book not found"}`
    assert.JSONEq(t, expectedNotFound, w.Body.String())
}
