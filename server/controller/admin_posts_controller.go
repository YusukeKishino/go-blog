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
	if err := c.db.Preload("Tags").First(&post, id).Error; err != nil {
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
	tagNames := ctx.PostFormArray("tags[]")

	var post model.Post
	if err := c.db.Preload("Tags").First(&post, id).Error; err != nil {
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

	err := c.db.Transaction(func(db *gorm.DB) error {
		if err := c.deleteTags(db, &post, tagNames); err != nil {
			return err
		}
		if err := c.createNewTags(db, tagNames, &post); err != nil {
			return err
		}

		if err := c.db.Save(&post).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/admin/posts/show/%d", post.ID))
}

func (c *AdminPostsController) deleteTags(db *gorm.DB, post *model.Post, tagNames []string) error {
	var newTags []model.Tag
	var deleteTags []model.Tag
	for _, tag := range post.Tags {
		found := false
		for _, tagName := range tagNames {
			if tag.Name == tagName {
				newTags = append(newTags, tag)
				found = true
				break
			}
		}
		if !found {
			deleteTags = append(deleteTags, tag)
		}
	}
	post.Tags = newTags
	if err := db.Model(&post).Association("Tags").Delete(deleteTags); err != nil {
		return err
	}
	return nil
}

func (c *AdminPostsController) createNewTags(db *gorm.DB, tagNames []string, post *model.Post) error {
	var tags, newTags []model.Tag
	if err := db.Where("name IN ?", tagNames).Find(&tags).Error; err != nil {
		return err
	}
	for _, tagName := range tagNames {
		found := false
		for _, tag := range tags {
			if tagName == tag.Name {
				found = true
				newTags = append(newTags, tag)
				break
			}
		}
		if !found {
			newTags = append(newTags, model.Tag{
				Name: tagName,
			})
		}
	}
	post.Tags = newTags
	return nil
}
