package controller

import (
	"net/http"
	"rest-api-template-go/dto"
	"rest-api-template-go/service"
	"rest-api-template-go/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	res, err := service.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

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

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := service.GetBook(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}
func UpdateBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func DeleteBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
