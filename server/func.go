package server

import (
	"html/template"
	"regexp"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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
