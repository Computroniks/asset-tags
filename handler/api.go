// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Computroniks/asset-tags/util"
)

func GetTag(w http.ResponseWriter, req *http.Request) (int, error) {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		return http.StatusBadRequest, errors.New("400 bad request")
	} else {
		tag, err := util.DB.GetTag(prefix)

		if err == sql.ErrNoRows {
			return http.StatusNotFound, err
		}

		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"tag\": \"%s\"}", tag)))
		return http.StatusOK, nil
	}
}

func IncrementTag(w http.ResponseWriter, req *http.Request) (int, error) {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		return http.StatusBadRequest, errors.New("400 bad request")
	} else {
		err := util.DB.IncrementTag(prefix)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte("{\"status\": 200}"))
		return http.StatusOK, nil
	}
}

func GetPrefixes(w http.ResponseWriter, req *http.Request) (int, error) {
	prefixes, err := util.DB.GetPrefixes()

	if err == sql.ErrNoRows {
		return http.StatusNotFound, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	prefixJson := ""
	for _, prefix := range prefixes {
		prefixJson += "\"" + prefix + "\","
	}
	if len(prefixJson) > 1 {
		prefixJson = prefixJson[:len(prefixJson)-1]
	}

	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"prefixes\":[%s]}", prefixJson)))
	return http.StatusOK, nil
}

func AddPrefix(w http.ResponseWriter, req *http.Request) (int, error) {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		return http.StatusBadRequest, errors.New("400 bad request")
	} else {
		err := util.DB.AddPrefix(prefix)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte("{\"status\": 200}"))
		return http.StatusOK, nil
	}
}
