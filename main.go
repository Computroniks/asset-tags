// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strconv"

	"github.com/Computroniks/asset-tags/router"
	"github.com/Computroniks/asset-tags/store/mysql"
	"github.com/Computroniks/asset-tags/util"
)

func init() {
	log.Println("Fetching environment variables")
	util.BindAddr = util.Getenv(util.BindAddrEnv, util.DefaultBindAddr)
	util.DatabaseAddr = util.Mustgetenv(util.DatabaseAddrEnv)
	util.DatabaseUsr = util.Mustgetenv(util.DatabaseUsrEnv)
	util.DatabasePwd = util.Mustgetenv(util.DatabasePwdEnv)
	util.DatabaseName = util.Mustgetenv(util.DatabaseNameEnv)

	var err error
	util.TagLength, err = strconv.Atoi(
		util.Getenv(
			util.TagLengthEnv,
			strconv.Itoa(util.DefaultTagLength),
			),
		)

	if err != nil {
		log.Println("Failed to read tag length:", err)
		log.Println("Using default tag length of", util.DefaultTagLength)
		util.TagLength = util.DefaultTagLength
	}
}

func main() {
	db, err := mysql.New(
		util.DatabaseAddr,
		util.DatabaseUsr,
		util.DatabasePwd,
		util.DatabaseName,
	)

	if err != nil {
		log.Fatalln(err)
		
	}

	router.RegisterHandlers(db)
	router.Start(util.BindAddr)
}