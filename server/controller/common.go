package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/YusukeKishino/go-blog/server/middleware"
)

func h(ctx *gin.Context, h2 gin.H) gin.H {
	fmt.Println(ctx.GetString(middleware.AdminKey))
	return merge(gin.H{
		"admin": ctx.GetString(middleware.AdminKey) != "",
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
