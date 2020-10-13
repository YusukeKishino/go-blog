package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/YusukeKishino/go-blog/client"
)

type AdminImagesController struct {
	s3Client client.S3Client
}

func NewAdminImagesController(s3Client client.S3Client) *AdminImagesController {
	return &AdminImagesController{s3Client: s3Client}
}

func (c *AdminImagesController) Upload(ctx *gin.Context) {
	image, err := ctx.FormFile("image")
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ext := filepath.Ext(image.Filename)
	filename := uuid.New().String() + ext

	f, err := image.Open()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	defer f.Close()

	location, err := c.s3Client.Upload(f, filename)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"location": location,
	})
}
