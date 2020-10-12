package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

	ctx.HTML(http.StatusOK, "posts_index.html.tmpl", h(ctx, gin.H{
		"posts": posts,
	}))
}

func (c *PostsController) Show(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.Wrap(err, fmt.Sprintf("failed to convert %s to int", idParam)))
		return
	}

	var post model.Post
	if err := c.db.Scopes(published).First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "posts_show.html.tmpl", h(ctx, gin.H{
		"post": post,
	}))
}

func published(db *gorm.DB) *gorm.DB {
	return db.Where("posts.status = ?", model.Published)
}
