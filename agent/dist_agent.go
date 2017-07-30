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

var conf *config.Config
var logger *log.Logger

func init()  {
	conf = libs.InitConfig("config.json")
	logger = libs.InitLogger("agent.log")
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
	logger.Notice("DIST Agent Module")
	port := conf.GetString("port","8080")
	http.ListenAndServe(":"+port, nil)
}
