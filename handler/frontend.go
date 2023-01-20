// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package handler

import (
	"log"
	"net/http"

	"github.com/Computroniks/asset-tags/templates"
	"github.com/Computroniks/asset-tags/util"
)

type templateData struct {
	Current string
	Prefixes []string
	CurrentPrefix string
}

func Index(w http.ResponseWriter, req *http.Request) int {
	prefix := req.URL.Query().Get("prefix")

	prefixes, err := util.DB.GetPrefixes()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500"))
		return http.StatusInternalServerError
	}

	if prefix == "" {
		if len(prefixes) == 0 {
			http.Redirect(w, req, "/settings", http.StatusSeeOther)
			return http.StatusSeeOther
		}
		prefix = prefixes[0]
	} 

	var tag string
	tag, err = util.DB.GetTag(prefix)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500"))
		return http.StatusInternalServerError
	}
	
	templates.Templates["index"].Execute(w, templateData{Current: tag, Prefixes: prefixes, CurrentPrefix: prefix})
	return http.StatusOK
}
