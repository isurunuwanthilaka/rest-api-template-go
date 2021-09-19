package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuthors godoc
// @Summary Show all authors
// @Description This will return list of authors in the system
// @Tags authors
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.AuthorResDto
// @Failure 400 {object} errors.RestError
// @Failure 404 {object} errors.RestError
// @Failure 500 {object} errors.RestError
// @Router /authors [get]
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
