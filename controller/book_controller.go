package controller

import (
	"net/http"
	"rest-api-template-go/dto"
	"rest-api-template-go/service"
	"rest-api-template-go/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks Show all books
// @Summary Show all books
// @Description This will return list of books in the system
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.BookResDto
// @Failure 400 {object} errors.RestError
// @Router /books [get]
func GetBooks(c *gin.Context) {
	res, err := service.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// AddBook Add a book
// @Summary Add a book
// @Description This will add a book to the system
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.BookResDto
// @Failure 400 {object} errors.RestError
// @Router /books [post]
func AddBook(c *gin.Context) {
	var newBook dto.BookReqDto
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	res, err := service.AddBook(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// GetBook Get a book by id
// @Summary Get a book by id
// @Description This will return a book in the system given the book id
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.BookResDto
// @Failure 400 {object} errors.RestError
// @Router /books/:id [post]
func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := service.GetBook(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// UpdateBook Update a book by id
// @Summary Update a book by id
// @Description This will update a book in the system given the book id
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.BookResDto
// @Failure 400 {object} errors.RestError
// @Router /books/:id [put]
func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var newBook dto.BookReqDto
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	res, err := service.UpdateBook(newBook, uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// DeleteBook Delete a book by id
// @Summary Delete a book by id
// @Description This will soft delete a book in the system given the book id
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} errors.RestError
// @Router /books/:id [delete]
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteBook(uint(id)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "Deleted successfully.")
}
