/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Control Module Drivered by Golang
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"fmt"
	"net/http"
	"DIST/libs"
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
)

var c_conf *config.Config
var c_logger *log.Logger

func init() {
	c_conf = libs.InitConfig("config.json")
	c_logger = libs.InitLogger("control.log")
	initRouter()
}

func initRouter() {
	http.HandleFunc("/stat", getSlaveStat)
}

func main() {
	fmt.Println("DIST Control Module")
	c_logger.Notice("DIST Control Module")
	port := c_conf.GetString("port","8090")
	http.ListenAndServe(":"+port, nil)
}
