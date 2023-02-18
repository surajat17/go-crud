# go-crud

This is a simple REST API for managing a collection of books using the Gin framework.
The API supports the following operations:

- Create a book: POST /books
- Read a book: GET /books/:id
- Update a book: PUT /books/:id
- Delete a book: DELETE /books/:id

Each book has an ID (integer), a title (string), an author (string), and a published date (string).

To run the API on your local machine, you need to install Go and the Gin framework.
Follow these steps:

1. Install Go: https://golang.org/doc/install
2. Install Gin: `go get -u github.com/gin-gonic/gin`
3. Save this file as `main.go`
4. In the terminal, navigate to the directory where you saved `main.go`
5. Run the command `go run main.go`
6. The API is now running on http://localhost:8080

To test the API, you can use the following curl commands:

Create a book:
"curl -X POST -H "Content-Type: application/json" -d '{"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "published": "April 10, 1925"}' http://localhost:8080/books"

Read a book:
"curl http://localhost:8080/books/1"

Update a book:
curl -X PUT -H "Content-Type: application/json" -d '{"title": "The Great Gatsby (Updated)", "author": "F. Scott Fitzgerald (Updated)", "published": "April 10, 1925"}' http://localhost:8080/books/1

Delete a book:
curl -X DELETE http://localhost:8080/books/1

