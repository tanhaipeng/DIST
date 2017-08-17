/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Libs and Global values
 * 2017/7/25, by Tan Haipeng, create
 */

package libs

import "net/http"
import (
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
	"io"
)

func InitConfig(path string) *config.Config {
	conf := config.New()
	conf.Load(path)
	return conf
}

func InitLogger(path string) *log.Logger {
	logger := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = path
	t2.MaxLevel = log.LevelDebug
	logger.Targets = append(logger.Targets, t1, t2)
	logger.Open()
	return logger
}

func GetRequest(req *http.Request, key string) string {
	// get
	query := req.URL.Query()
	for k, v := range query {
		if len(v) > 0 {
			if k == key {
				return v[0]
			}
		}
	}
	// post
	req.ParseForm()
	if len(req.Form[key]) > 0 {
		return req.Form[key][0]
	}
	return ""
}

func SendResponse(rsp http.ResponseWriter, ret string) {
	io.WriteString(rsp, ret)
}