// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package handler

import (
	"net/http"

	"github.com/Computroniks/asset-tags/templates"
	"github.com/Computroniks/asset-tags/util"
)

type templateData struct {
	Current string
	Prefixes []string
	CurrentPrefix string
}

func Index(w http.ResponseWriter, req *http.Request) (int, error) {
	prefix := req.URL.Query().Get("prefix")

	prefixes, err := util.DB.GetPrefixes()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if prefix == "" {
		if len(prefixes) == 0 {
			http.Redirect(w, req, "/settings", http.StatusSeeOther)
			return http.StatusSeeOther, nil
		}
		prefix = prefixes[0]
	} 

	var tag string
	tag, err = util.DB.GetTag(prefix)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	templates.Templates["index"].Execute(w, templateData{Current: tag, Prefixes: prefixes, CurrentPrefix: prefix})
	return http.StatusOK, nil
}
