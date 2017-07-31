/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Module Drivered by Golang
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import "fmt"
import "net/http"
import (
	"DIST/libs"
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
)

var a_conf *config.Config
var a_logger *log.Logger

func init() {
	a_conf = libs.InitConfig("agent/config.json")
	a_logger = libs.InitLogger("agent/agent.log")
	initRouter()
}

func initRouter() {
	http.HandleFunc("/stat", getSelfStat)
	http.HandleFunc("/start", startSlave)
	http.HandleFunc("/stop", stopSlave)
	http.HandleFunc("/update", updateSlave)
}

func main() {
	fmt.Println("DIST Agent Module")
	a_logger.Notice("DIST Agent Module")
	port := a_conf.GetString("port", "8020")
	http.ListenAndServe(":"+port, nil)
}
