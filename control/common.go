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
	"io"
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

func getRequest(req *http.Request, rtype string) map[string]string {
	var params = make(map[string]string)
	if rtype == "get" {
		query := req.URL.Query()
		for k, v := range query {
			if len(v) > 0 {
				params[k] = v[0]
			}
		}
	}
	if rtype == "post" {

	}
	return params
}

func sendResponse(rsp http.ResponseWriter, rtype string) {
	if rtype == "json" {
		io.WriteString(rsp, combineRetData(0, "success"))
	} else {
		io.WriteString(rsp, combineRetData(100, "type error"))
	}
}

func combineRetData(code int, msg string) string {

	return ""
}

func callSlaveApi(url string, params map[string]string) bool {
	return true
}
