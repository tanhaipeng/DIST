/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Router handler
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"net/http"
	"fmt"
	"../libs"
)

func getSelfStat(rsp http.ResponseWriter, req *http.Request) {
	getData := libs.GetRequest(req, "get")
	postData := libs.GetRequest(req, "post")
	fmt.Println(getData, postData)
	libs.SendResponse(rsp, "json")
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
