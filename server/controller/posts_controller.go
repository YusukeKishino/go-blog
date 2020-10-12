package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/model"
)

type PostsController struct {
	db *gorm.DB
}

func NewPostsController(db *gorm.DB) *PostsController {
	return &PostsController{db: db}
}

func (c *PostsController) Index(ctx *gin.Context) {
	var posts []*model.Post
	if err := c.db.Scopes(published).Order("published_at DESC").Find(&posts).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "posts_index.html.tmpl", gin.H{
		"posts": posts,
	})
}

func (c *PostsController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var post model.Post
	if err := c.db.Scopes(published).First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "posts_show.html.tmpl", gin.H{
		"post": post,
	})
}

func published(db *gorm.DB) *gorm.DB {
	return db.Where("posts.status = ?", model.Published)
}
