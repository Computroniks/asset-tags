// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package handler

import (
	"net/http"

	"github.com/Computroniks/asset-tags/util"
)

func HTTP400(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	util.Views.Render(w, "400.html", nil)
}

func HTTP404(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	util.Views.Render(w, "404.html", nil)
}

func HTTP500(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	util.Views.Render(w, "500.html", nil)
}
