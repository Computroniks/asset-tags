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
}

func Index(w http.ResponseWriter, req *http.Request) int {
	prefix := req.URL.Query().Get("prefix")
	var tag string
	if prefix == "" {
		// TODO: Get first tag
	} else {
		var err error
		tag, err = util.DB.GetTag(prefix)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
			return http.StatusInternalServerError
		}
	}
	templates.Templates["index"].Execute(w, templateData{Current: tag})
	return http.StatusOK
}
