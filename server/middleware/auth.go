package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	AdminKey = "admin"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	admin := session.Get(AdminKey)

	if admin == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
