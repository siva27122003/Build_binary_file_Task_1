package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
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

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func main() {
    r := gin.Default()
    r.GET("/books", getBooks)
    r.Run(":8081") // Run server on port 8081
}