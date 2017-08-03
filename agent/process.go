/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Porcess Drivered by Golang
 * 2017/8/2, by Tan Haipeng, create
 */

package main

import (
	"encoding/json"
	"os/exec"
	"strings"
	"fmt"
)

func FixRetData(code int, msg string, data string) string {
	var ret RetType
	ret.Code = code
	ret.Msg = msg
	ret.Data = data
	str, _ := json.Marshal(ret)
	return string(str)
}

func GetSysInfo() (string, string) {
	cmd := "ps aux | grep login | grep -v grep | awk '{print $2}'"
	res, err := exec.Command("sh", "-c", cmd).Output()
	if err == nil {
		pids := strings.Split(strings.Trim(string(res), "\n"), "\n")
		if len(pids) > 0 {
			cmd = fmt.Sprintf("ps -p %s -o lstart", pids[0])
			res, err = exec.Command("sh", "-c", cmd).Output()
			if err == nil {
				info := strings.Split(strings.Trim(string(res), "\n"), "\n")
				if len(info) > 1 {
					return info[0], info[1]
				}
			}
		}
	}
	return "", ""
}
