/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Porcess Drivered by Golang
 * 2017/8/2, by Tan Haipeng, create
 */

package main

import "encoding/json"

func FixRetData(code int, msg string, data string) string {
	var ret RetType
	ret.Code = code
	ret.Msg = msg
	ret.Data = data
	str, _ := json.Marshal(ret)
	return string(str)
}
