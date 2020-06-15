package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ping service
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
