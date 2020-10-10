package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/YusukeKishino/go-blog/config"
	"github.com/YusukeKishino/go-blog/server/controller"
	"github.com/YusukeKishino/go-blog/server/middleware"
)

type Controllers struct {
	dig.In
	Index      *controller.IndexController
	AdminPosts *controller.AdminPostsController
	Posts      *controller.PostsController
	Login      *controller.LoginController
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
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html.tmpl", gin.H{})
	})
	engine.Static("/images", "./server/assets/src/images")
	engine.GET("/", r.controllers.Index.Index)
	posts := engine.Group("/posts")
	{
		posts.GET("/", r.controllers.Posts.Index)
		posts.GET("/show/:id", r.controllers.Posts.Show)
	}

	engine.GET("/login", r.controllers.Login.Show)
	engine.POST("/login", r.controllers.Login.Login)

	adminGroup := engine.Group("/admin", middleware.AuthRequired)
	{
		posts := adminGroup.Group("/posts")
		{
			posts.GET("/", r.controllers.AdminPosts.Index)
			posts.GET("/new", r.controllers.AdminPosts.New)
			posts.GET("/show/:id", r.controllers.AdminPosts.Show)
			posts.GET("/edit/:id", r.controllers.AdminPosts.Edit)
			posts.POST("/:id", r.controllers.AdminPosts.Update)
		}
	}
}
