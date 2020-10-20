package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/config"
	"github.com/YusukeKishino/go-blog/server/middleware"
)

func h(ctx *gin.Context, db *gorm.DB, h2 gin.H) gin.H {
	return merge(gin.H{
		"admin":     ctx.GetString(middleware.AdminKey) != "",
		"isProd":    config.IsProd(),
		"tagCounts": tagCounts(db),
	}, h2)
}

func merge(m ...gin.H) gin.H {
	ans := make(gin.H)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}

	}
	return ans
}

type TagCount struct {
	Name  string
	Count int
}

func tagCounts(db *gorm.DB) []TagCount {
	var results []TagCount
	// ここでエラーハンドリングをしてもあまり嬉しくないので握りつぶす
	_ = db.Raw(`
SELECT 
	tags.name as 'name',
	count(1) as 'count'
FROM tags 
INNER JOIN post_tags 
	ON tags.id = post_tags.tag_id
INNER JOIN posts
	on posts.id = post_tags.post_id and posts.status = 'published'
GROUP BY tags.id`).Scan(&results)
	return results
}
