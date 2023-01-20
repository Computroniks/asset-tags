// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package util

import (
	"github.com/Computroniks/asset-tags/store"
)

// Runtime config
var (
	BindAddr string
	DatabaseAddr string
	DatabaseUsr string
	DatabasePwd string
	DatabaseName string
	DatabaseTimeout string
	TagLength int
	BasePath string
	DB store.Store
)

// Environment variables
const (
	BindAddrEnv = "BIND_ADDRESS"
	DatabaseAddrEnv = "DB_ADDRESS"
	DatabaseUsrEnv = "DB_USER"
	DatabasePwdEnv = "DB_PASSWORD"
	DatabaseNameEnv = "DB_NAME"
	DatabaseTimeoutEnv = "DB_TIMEOUT"
	TagLengthEnv = "TAG_LENGTH"
	BasePathEnv = "BASE_PATH"
)

// Defaults
const (
	DefaultBindAddr = ":3000"
	DefaultTagLength = 6
	DefaultBasePath = "/"
	DefaultDBTimeout = "5"
)
