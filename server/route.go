package server

import (
	"github.com/gin-gonic/gin"

	"github.com/YusukeKishino/go-blog/config"
)

func setRoutes(engine *gin.Engine) {
	if config.IsProd() {
		engine.GET("/webpack/*name", func(c *gin.Context) {
			c.File("server/assets/public/webpack/" + c.Param("name"))
		})
	}
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html.tmpl", gin.H{})
	})
}
