package controller

import (
	"fmt"
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
	if err := c.db.Order("id DESC").Where("status = ?", model.Published).Find(&posts).Error; err != nil {
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
	if err := c.db.First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "posts_show.html.tmpl", gin.H{
		"post": post,
	})
}

func (c *PostsController) New(ctx *gin.Context) {
	post := model.Post{
		Title:   "タイトル",
		Content: "",
		Status:  model.Draft,
	}
	if err := c.db.Create(&post).Error; err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/admin/posts/edit/%d", post.ID))
}
