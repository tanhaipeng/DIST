/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Control Router handler
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"net/http"
)

func getSlaveStat(rsp http.ResponseWriter, req *http.Request) {
	getRequest(req, "get")
	sendResponse(rsp, "json")
}
