package server

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/go-webpack/webpack"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	// TODO: Fixme
	"github.com/YusukeKishino/go-blog/config"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	setupWebpack()

	engine := gin.New()
	engine.Use(ginlogrus.Logger(logrus.StandardLogger()), gin.Recovery())
	engine.HTMLRender = loadTemplates("server/views")

	setRoutes(engine)

	return &Server{Engine: engine}
}

func setupWebpack() {
	webpack.DevHost = "localhost:3808" // default
	webpack.Plugin = "manifest"        // defaults to stats for compatibility
	webpack.FsPath = "./server/assets/public/webpack"
	webpack.Init(config.IsDev())
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html.tmpl")
	if err != nil {
		panic(err.Error())
	}

	pages, err := filepath.Glob(templatesDir + "/pages/*.html.tmpl")
	if err != nil {
		panic(err.Error())
	}

	funcMap := template.FuncMap{
		"asset": webpack.AssetHelper,
	}

	// Generate our templates map from our layouts/ and pages/ directories
	for _, include := range pages {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
	}
	return r
}
