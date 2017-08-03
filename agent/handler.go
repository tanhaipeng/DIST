/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Router handler
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"net/http"
	"fmt"
	"DIST/libs"
	"encoding/json"
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
	getData := libs.GetRequest(req, "get")
	postData := libs.GetRequest(req, "post")
	fmt.Println(getData, postData)
	libs.SendResponse(rsp, "json")
}

func updateSlave(rsp http.ResponseWriter, req *http.Request) {
	getData := libs.GetRequest(req, "get")
	postData := libs.GetRequest(req, "post")
	fmt.Println(getData, postData)
	libs.SendResponse(rsp, "json")
}

func stopSlave(rsp http.ResponseWriter, req *http.Request) {
	getData := libs.GetRequest(req, "get")
	postData := libs.GetRequest(req, "post")
	fmt.Println(getData, postData)
	libs.SendResponse(rsp, "json")
}
