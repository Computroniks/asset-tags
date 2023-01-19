// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Computroniks/asset-tags/util"
)

type Router struct {
	routes map[string]map[string]func(http.ResponseWriter, *http.Request)int
	errors map[int]func(http.ResponseWriter, *http.Request)
	basePath string
}

// Return pointer to router instance
func New(basePath string) *Router {
	rtr := Router{
		basePath: basePath,
		routes: make(map[string]map[string]func(http.ResponseWriter, *http.Request)int),
		errors: make(map[int]func(http.ResponseWriter, *http.Request)),
	}
	return &rtr
}

// Add GET handler
func (o *Router) GET(path string, handler func(http.ResponseWriter, *http.Request)int) {
	if _, exists := o.routes[path]; !exists {
		o.routes[path] = make(map[string]func(http.ResponseWriter, *http.Request)int)
	}
	o.routes[path]["GET"] = handler
}

// Add POST handler
func (o *Router) POST(path string, handler func(http.ResponseWriter, *http.Request)int) {
	o.routes[path]["POST"] = handler
}

// Start the server
func (o *Router) Start(addr string) {
	log.Println("Starting server on", addr)
	http.HandleFunc(o.basePath, o.routeRequest)
	http.ListenAndServe(util.BindAddr, nil)
}

func (o *Router) sendError(w http.ResponseWriter, req *http.Request, status int) {
	handler, exists := o.errors[status]
	log.Println(req.Method, req.URL, req.RemoteAddr, status)
	if exists {
		handler(w, req)
	} else {
		w.WriteHeader(status)
		w.Write([]byte(strconv.Itoa(status)))
	}

}

// Route the users request to the correct handler
func (o *Router) routeRequest(w http.ResponseWriter, req *http.Request) {
	URL := req.URL
	method := req.Method
	remoteAddr := req.RemoteAddr

	handler, exists := o.routes[req.URL.Path][req.Method]
	if exists {
		status := handler(w, req)
		log.Println(method, URL, remoteAddr, status)
	} else {
		if _, exists := o.routes[req.URL.Path]; exists {
			o.sendError(w, req, http.StatusMethodNotAllowed)
		}
		o.sendError(w, req, http.StatusNotFound)
	}
}
