/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Router handler
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"net/http"
	"DIST/libs"
	"encoding/json"
	"os"
)

func getSelfStat(rsp http.ResponseWriter, req *http.Request) {
	item := libs.GetRequest(req, "item")
	ret := ""
	switch item {
	case "slave":
		var statRet StatType
		stat, time := GetSysInfo()
		statRet.Code = 0
		statRet.Msg = ""
		statRet.Data.Stat = stat
		statRet.Data.Time = time
		tmp, _ := json.Marshal(statRet)
		ret = string(tmp)
	default:
		ret = FixRetData(100, "item error", "")
	}
	if item == "" {
		ret = FixRetData(101, "item lost", "")
	}
	libs.SendResponse(rsp, ret)
}

func startSlave(rsp http.ResponseWriter, req *http.Request) {
	signal = true
	ExecTask()
	libs.SendResponse(rsp, FixRetData(0, "test start", ""))
}

func updateSlave(rsp http.ResponseWriter, req *http.Request) {
	task := libs.GetRequest(req, "task")
	if task != "" {
		file, err := os.OpenFile("agent/task", os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err == nil {
			defer file.Close()
			_, err = file.WriteString(task)
			if err != nil {
				libs.SendResponse(rsp, FixRetData(103, "file write err", ""))
			} else {
				libs.SendResponse(rsp, FixRetData(0, "task update", ""))
			}
		} else {
			libs.SendResponse(rsp, FixRetData(103, "file open err", ""))
		}
	} else {
		libs.SendResponse(rsp, FixRetData(102, "task empty", ""))
	}
}

func stopSlave(rsp http.ResponseWriter, req *http.Request) {
	signal = false
	libs.SendResponse(rsp, FixRetData(0, "test stop", ""))
}
