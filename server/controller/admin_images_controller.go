package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AdminImagesController struct {
}

func NewAdminImagesController() *AdminImagesController {
	return &AdminImagesController{}
}

func (c *AdminImagesController) Upload(ctx *gin.Context) {
	f, err := ctx.FormFile("image")
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Println(f)
}
