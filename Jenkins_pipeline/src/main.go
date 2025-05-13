package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "Book One", Author: "Author One"},
    {ID: "2", Title: "Book Two", Author: "Author Two"},
}

// Fetch all books
func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

// Fetch a book by ID
func getBookById(c *gin.Context) {
    id := c.Param("id")

    for _, book := range books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
    r := gin.Default()
    r.GET("/books", getBooks)
    r.GET("/books/:id", getBookById) // New route to get a book by ID
    r.Run(":8081") // Run server on port 8081
}
