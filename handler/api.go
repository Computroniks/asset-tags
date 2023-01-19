package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Computroniks/asset-tags/util"
)

func GetTag(w http.ResponseWriter, req *http.Request) int {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400"))
		return http.StatusBadRequest
	} else {
		tag, err := util.DB.GetTag(prefix)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
			return http.StatusInternalServerError
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"tag\": \"%s\"}", tag)))
		return http.StatusOK
	}
}

func IncrementTag(w http.ResponseWriter, req *http.Request) int {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400"))
		return http.StatusBadRequest
	} else {
		err := util.DB.IncrementTag(prefix)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
			return http.StatusInternalServerError
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte("{\"status\": 200}"))
		return http.StatusOK
	}
}

func GetPrefixes(w http.ResponseWriter, req *http.Request) int {
	prefixes, err := util.DB.GetPrefixes()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500"))
		return http.StatusInternalServerError
	}

	prefixJson := ""
	for _, prefix := range prefixes {
		prefixJson += "\"" + prefix + "\","
	}
	prefixJson = prefixJson[:len(prefixJson)-1]

	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"prefixes\":[%s]}", prefixJson)))
	return http.StatusOK
}

func AddPrefix(w http.ResponseWriter, req *http.Request) int {
	prefix := req.URL.Query().Get("prefix")
	if prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400"))
		return http.StatusBadRequest
	} else {
		err := util.DB.AddPrefix(prefix)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
			return http.StatusInternalServerError
		}

		w.Header().Add("Content-type", "application/json")
		w.Write([]byte("{\"status\": 200}"))
		return http.StatusOK
	}
}
