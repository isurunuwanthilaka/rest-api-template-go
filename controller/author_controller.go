package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetAuthor(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func AddAuthor(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func UpdateAuthor(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func DeleteAuthor(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
