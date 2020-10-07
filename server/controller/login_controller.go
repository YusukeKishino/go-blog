package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/model"
	"github.com/YusukeKishino/go-blog/server/middleware"
)

type LoginController struct {
	db *gorm.DB
}

func NewLoginController(db *gorm.DB) *LoginController {
	return &LoginController{db: db}
}

func (c *LoginController) Show(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html.tmpl", gin.H{})
}

func (c *LoginController) Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	name := ctx.PostForm("name")
	formPassword := ctx.PostForm("password")

	var admin model.Admin
	if err := c.db.First(&admin, model.Admin{Name: name}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.HTML(http.StatusUnauthorized, "login.html.tmpl", gin.H{})
			return
		}
		_ = ctx.Error(err)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(formPassword)); err != nil {
		ctx.HTML(http.StatusUnauthorized, "login.html.tmpl", gin.H{})
		return
	} else {
		session.Set(middleware.AdminKey, name)
		if err := session.Save(); err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.Redirect(http.StatusFound, "/admin/posts")
	}
}
