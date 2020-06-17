package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping ping service
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
