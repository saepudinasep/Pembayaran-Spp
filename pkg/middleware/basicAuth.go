package middleware

import (
	"Pembayaran-Spp/model/json"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		json.NewResponseUnauthorized(c, "Invalid Token", "00", "00")
		c.Abort()
		return
	}

	if user != "admin" || password != "admin" {
		json.NewResponseUnauthorized(c, "Unauthorized", "00", "00")
		c.Abort()
		return
	}
	c.Next()
}
