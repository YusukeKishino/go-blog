package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/model"
)

type IndexController struct {
	db *gorm.DB
}

func NewIndexController(db *gorm.DB) *IndexController {
	db = db.Scopes(published)
	return &IndexController{db: db}
}

func (c *IndexController) Index(ctx *gin.Context) {
	var posts []*model.Post
	if err := c.db.Order("id DESC").Find(&posts).Error; err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html.tmpl", h(ctx, gin.H{
		"posts": posts,
	}))
}
