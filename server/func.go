package server

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/YusukeKishino/go-blog/model"
)

var regxNewline = regexp.MustCompile(`\r\n|\r|\n`)

func MD2HTML(md string) template.HTML {
	replaced := regxNewline.ReplaceAllString(md, "\n")
	return template.HTML(markdown.ToHTML([]byte(replaced), getParser(), renderer()))
}

func getParser() *parser.Parser {
	return parser.NewWithExtensions(
		parser.CommonExtensions | parser.HardLineBreak,
	)
}

func renderer() markdown.Renderer {
	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	}
	return html.NewRenderer(opts)
}

func MaxLength(str string, length int) template.HTML {
	if len(strings.Split(str, "")) >= length {
		str = string([]rune(str)[:length]) + "..."
	}
	return template.HTML(str)
}

func ToPost(i interface{}) (*model.Post, error) {
	p, ok := i.(model.Post)
	if !ok {
		return &model.Post{}, fmt.Errorf("failed to vonvert")
	}
	return &p, nil
}
