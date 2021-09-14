package app

import (
	"rest-api-template-go/controller"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication start the server
func StartApplication() {

	// book apis
	router.GET(v1("/books"), controller.GetBooks)
	router.GET(v1("/books/:id"), controller.GetBook)
	router.PUT(v1("/books/:id"), controller.UpdateBook)
	router.DELETE(v1("/books/:id"), controller.DeleteBook)
	router.POST(v1("/books"), controller.AddBook)

	// author apis
	router.GET(v1("/authors"), controller.GetAuthors)
	router.GET(v1("/authors/:id"), controller.GetAuthor)
	router.PUT(v1("/authors/:id"), controller.UpdateAuthor)
	router.DELETE(v1("/authors/:id"), controller.DeleteAuthor)
	router.POST(v1("/authors"), controller.AddAuthor)

	router.Run(":8080")
}

func v1(path string) string {
	prepend := "/api/v1"
	return prepend + path
}
