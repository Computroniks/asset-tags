package templates

import (
	"embed"
	"html/template"
)

//go:embed *
var templateFiles embed.FS


var Templates map[string]*template.Template

func Init() {
	Templates = make(map[string]*template.Template)
	Templates["index"] = template.Must(template.ParseFS(templateFiles, "index.html"))
}