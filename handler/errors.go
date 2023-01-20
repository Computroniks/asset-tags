// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package handler

import (
	"net/http"

	"github.com/Computroniks/asset-tags/templates"
)

func HTTP400(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	templates.Templates["400"].Execute(w, nil)
}

func HTTP404(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.Templates["404"].Execute(w, nil)
}

func HTTP500(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.Templates["500"].Execute(w, nil)
}
