// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

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
	Templates["400"] = template.Must(template.ParseFS(templateFiles, "400.html"))
	Templates["404"] = template.Must(template.ParseFS(templateFiles, "404.html"))
	Templates["500"] = template.Must(template.ParseFS(templateFiles, "500.html"))
}