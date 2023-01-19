// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package util

// Runtime config
var (
	BindAddr string
	DatabaseAddr string
	DatabaseUsr string
	DatabasePwd string
	DatabaseName string
	TagLength int
)

// Environment variables
const (
	BindAddrEnv = "BIND_ADDRESS"
	DatabaseAddrEnv = "DB_ADDRESS"
	DatabaseUsrEnv = "DB_USER"
	DatabasePwdEnv = "DB_PASSWORD"
	DatabaseNameEnv = "DB_NAME"
	TagLengthEnv = "TAG_LENGTH"
)

// Defaults
const (
	DefaultBindAddr = ":3000"
	DefaultTagLength = 6
)
