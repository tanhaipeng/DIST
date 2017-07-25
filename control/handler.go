/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Control Router handler
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"net/http"
	"fmt"
)

func getSlaveStat(rsp http.ResponseWriter, req *http.Request) {
	getData := getRequest(req, "get")
	postData := getRequest(req, "post")
	fmt.Println(getData, postData)
	sendResponse(rsp, "json")
}
