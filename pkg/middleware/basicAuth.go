package middleware

import (
	"Pembayaran-Spp/model/json"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		json.NewResponseUnauthorized(c, "Invalid Token", "01", "01")
		c.Abort()
		return
	}

	if user != "admin" || password != "admin" {
		json.NewResponseUnauthorized(c, "Unauthorized", "01", "02")
		c.Abort()
		return
	}
	c.Next()
}
