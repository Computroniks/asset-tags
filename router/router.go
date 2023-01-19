// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package router

import (
	"log"

	"github.com/Computroniks/asset-tags/store"
)

// Create new router and register all endpoint handlers
func RegisterHandlers (store store.Store) {
	log.Println("Starting new router")
	// TODO: Add handlers
}

// Start the server
func Start(addr string) {
	log.Println("Starting server on", addr)
	// TODO: Actually start server
}