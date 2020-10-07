package server

import (
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func MD2HTML(md string) template.HTML {
	return template.HTML(markdown.ToHTML([]byte(md), nil, renderer()))
}

func renderer() markdown.Renderer {
	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	}
	return html.NewRenderer(opts)
}
