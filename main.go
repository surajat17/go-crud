package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Author    string `json:"author"`
    Published string `json:"published"`
}

var books = []Book{
    {ID: 1, Title: "To Kill a Mockingbird", Author: "Harper Lee", Published: "July 11, 1960"},
}

func main() {
    router := gin.Default()

    // Create
    router.POST("/books", func(c *gin.Context) {
        var book Book
        if err := c.ShouldBindJSON(&book); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        book.ID = len(books) + 1
        books = append(books, book)

        c.JSON(http.StatusCreated, book)
    })

    // Read
    router.GET("/books/:id", func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        for _, book := range books {
            if book.ID == id {
                c.JSON(http.StatusOK, book)
                return
            }
        }

        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
    })

    // Update
    router.PUT("/books/:id", func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        var newBook Book
        if err := c.ShouldBindJSON(&newBook); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        for i, book := range books {
            if book.ID == id {
                newBook.ID = book.ID
                books[i] = newBook

                c.JSON(http.StatusOK, newBook)
                return
            }
        }

        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
    })

    // Delete
    router.DELETE("/books/:id", func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        for i, book := range books {
            if book.ID == id {
                books = append(books[:i], books[i+1:]...)
                c.JSON(http.StatusNoContent, nil)
                return
            }
        }

        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
    })

    router.Run(":8080")
}