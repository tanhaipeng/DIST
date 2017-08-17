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
	"io/ioutil"
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
					return info[0], strings.Trim(info[1], " ")
				}
			}
		}
	}
	return "", ""
}

func GetTask() (TaskType, error) {
	var rTask TaskType
	task, err := ioutil.ReadFile("agent/task")
	if err == nil {
		if task != nil {
			err = json.Unmarshal([]byte(task), &rTask)
		}
	}
	return rTask, err
}

func HealthCheck() (host string, status bool) {
	return "127.0.0.1", true
}

func QueryString(data []FieldType) string {
	var ret = ""
	for _, elem := range data {
		if ret == "" {
			ret = elem.name + "=" + elem.value
		} else {
			ret = ret + "&" + elem.name + "=" + elem.value
		}
	}
	return ret
}

func execTask() {
	task, err := GetTask()
	if err != nil {
		a_logger.Error(err.Error())
		return
	}
	if strings.ToLower(task.Type) == "get" {
		qStr := QueryString(task.Field)
		qStr = task.Ip + ":" + task.Port + "?" + qStr
		fmt.Println(qStr)
	}
	if strings.ToLower(task.Type) == "post" {
		qByte, err := json.Marshal(task.Field)
		if err != nil {
			a_logger.Error(err.Error())
			return
		}
		var qStr = string(qByte)
		fmt.Println(qStr)
	}
}
