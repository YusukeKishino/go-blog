package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/YusukeKishino/go-blog/config"
	"github.com/YusukeKishino/go-blog/server/controller"
)

type Controllers struct {
	dig.In
	Posts *controller.PostsController
}

type Router struct {
	controllers Controllers
}

func NewRouter(controllers Controllers) *Router {
	return &Router{controllers: controllers}
}

func (r *Router) setRoutes(engine *gin.Engine) {
	if config.IsProd() {
		engine.GET("/webpack/*name", func(c *gin.Context) {
			c.File("server/assets/public/webpack/" + c.Param("name"))
		})
	}
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html.tmpl", gin.H{})
	})
	posts := engine.Group("/posts")
	{
		posts.GET("/", r.controllers.Posts.Index)
		posts.GET("/new", r.controllers.Posts.New)
		posts.GET("/show/:id", r.controllers.Posts.Show)
		posts.GET("/edit/:id", r.controllers.Posts.Edit)
		posts.POST("/:id", r.controllers.Posts.Update)
	}
}
