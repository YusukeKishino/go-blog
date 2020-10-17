package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/model"
)

type AdminPostsController struct {
	db *gorm.DB
}

func NewAdminPostsController(db *gorm.DB) *AdminPostsController {
	return &AdminPostsController{db: db}
}

func (c *AdminPostsController) Index(ctx *gin.Context) {
	var posts []*model.Post
	if err := c.db.Order("id DESC").Find(&posts).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "admin_posts_index.html.tmpl", h(ctx, gin.H{
		"posts": posts,
	}))
}

func (c *AdminPostsController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var post model.Post
	if err := c.db.First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "admin_posts_show.html.tmpl", h(ctx, gin.H{
		"post": post,
	}))
}

func (c *AdminPostsController) New(ctx *gin.Context) {
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

func (c *AdminPostsController) Edit(ctx *gin.Context) {
	id := ctx.Param("id")
	var post model.Post
	if err := c.db.Preload("Tags").First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}
	var tags []model.Tag
	if err := c.db.Order("name").Find(&tags).Error; err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.HTML(http.StatusOK, "admin_posts_edit.html.tmpl", h(ctx, gin.H{
		"post": post,
		"tags": tags,
	}))
}

func (c *AdminPostsController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.PostForm("title")
	status := ctx.PostForm("status")
	content := ctx.PostForm("content")
	var post model.Post
	if err := c.db.First(&post, id).Error; err != nil {
		_ = ctx.Error(err)
		return
	}
	post.Title = title
	post.Status = model.PostStatus(status)
	post.Content = content

	if post.IsPublished() && !post.PublishedAt.Valid {
		post.PublishedAt.Valid = true
		post.PublishedAt.Time = time.Now()
	}
	if err := c.db.Save(&post).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/admin/posts/show/%d", post.ID))
}
