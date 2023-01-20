// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *.html
var templateFiles embed.FS

type Template struct {
    templates *template.Template
}

func New() *Template {
    templates := template.Must(template.New("").ParseFS(templateFiles, "*.html"))
    return &Template{
        templates: templates,
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
    tmpl := template.Must(t.templates.Clone())
    tmpl = template.Must(tmpl.ParseFS(templateFiles, name))
    return tmpl.ExecuteTemplate(w, name, data)
}
