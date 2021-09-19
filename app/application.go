package app

import (
	"rest-api-template-go/controller"
	_ "rest-api-template-go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// StartApplication godoc
func StartApplication() {
	router := gin.Default()
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func v1(path string) string {
	prepend := "/api/v1"
	return prepend + path
}
