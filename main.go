// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strconv"

	"github.com/Computroniks/asset-tags/handler"
	"github.com/Computroniks/asset-tags/router"
	"github.com/Computroniks/asset-tags/store/mysql"
	"github.com/Computroniks/asset-tags/templates"
	"github.com/Computroniks/asset-tags/util"
)

func init() {
	log.Println("Fetching environment variables")
	util.BindAddr = util.Getenv(util.BindAddrEnv, util.DefaultBindAddr)
	util.DatabaseAddr = util.Mustgetenv(util.DatabaseAddrEnv)
	util.DatabaseUsr = util.Mustgetenv(util.DatabaseUsrEnv)
	util.DatabasePwd = util.Mustgetenv(util.DatabasePwdEnv)
	util.DatabaseName = util.Mustgetenv(util.DatabaseNameEnv)
	util.BasePath = util.Getenv(util.BasePathEnv, util.DefaultBasePath)

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

	log.Println("Initialising templates")
	templates.Init()
}

func main() {
	var err error
	util.DB, err = mysql.New(
		util.DatabaseAddr,
		util.DatabaseUsr,
		util.DatabasePwd,
		util.DatabaseName,
	)

	if err != nil {
		log.Fatalln(err)	
	}

	defer util.DB.Close()

	app := router.New(util.BasePath)
	app.GET("/", handler.Index)
	app.GET("/api/tag", handler.GetTag)
	app.POST("/api/tag", handler.IncrementTag)
	app.GET("/api/prefix", handler.GetPrefixes)
	app.POST("/api/prefix", handler.AddPrefix)
	app.Error(400, handler.HTTP400)
	app.Error(404, handler.HTTP404)
	app.Error(500, handler.HTTP500)
	app.Start(util.BindAddr)
}