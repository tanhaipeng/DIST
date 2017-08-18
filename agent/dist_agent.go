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
	"encoding/json"
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
	/* ==for test
	var task TaskType
	task.Ip = "127.0.0.1"
	task.Port = "8020"
	task.Method = "/update"
	task.Count = 10
	task.Time = 100
	task.Type = "get"
	task.Field = []FieldType{{"id", "12"}, {"age", "26"}}
	str, _ := json.Marshal(task)
	fmt.Println(string(str))
	*/
	ExecTask()
	fmt.Println("DIST Agent Module")
	a_logger.Notice("DIST Agent Module")
	port := a_conf.GetString("port", "8020")
	http.ListenAndServe(":"+port, nil)
}
