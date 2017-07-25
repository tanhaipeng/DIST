/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Control Libs and Global values
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import "net/http"
import (
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
)

var conf *config.Config
var logger *log.Logger

func initConfig() {
	conf = config.New()
	conf.Load("config.json")
}

func initLogger() {
	logger = log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = "control.log"
	t2.MaxLevel = log.LevelDebug
	logger.Targets = append(logger.Targets, t1, t2)
	logger.Open()
}

func getRequest(req *http.Request, rtype string) {

}

func sendResponse(rsp http.ResponseWriter, rtype string) {

}
