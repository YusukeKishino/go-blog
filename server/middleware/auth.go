package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	AdminKey = "admin"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	admin := session.Get(AdminKey)
	if admin != nil {
		c.Set(AdminKey, admin)
	}
	c.Next()
}

func AuthRequired(c *gin.Context) {
	if c.GetString(AdminKey) == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
