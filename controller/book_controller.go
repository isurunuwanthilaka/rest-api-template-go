package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func AddBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func UpdateBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func DeleteBook(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
