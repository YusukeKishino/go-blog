package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()
	if err != nil {
		cause := errors.Cause(err.Err)
		if errors.Is(cause, gorm.ErrRecordNotFound) {
			c.HTML(http.StatusNotFound, "404.html.tmpl", gin.H{})
		} else {
			c.HTML(http.StatusInternalServerError, "500.html.tmpl", gin.H{})
		}
	}
}
