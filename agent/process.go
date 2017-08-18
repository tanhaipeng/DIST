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

func QueryString(data []FieldType, rtype string) string {
	var ret = ""
	if rtype == "get" {
		for _, elem := range data {
			if ret == "" {
				ret = elem.Name + "=" + elem.Value
			} else {
				ret = ret + "&" + elem.Name + "=" + elem.Value
			}
		}
		return ret
	}
	if rtype == "post" {
		var pdata = make(map[string]string)
		for _, elem := range data {
			pdata[elem.Name] = elem.Value
		}
		rbyte, err := json.Marshal(pdata)
		if err == nil {
			ret = string(rbyte)
		}
	}
	return ret
}

func ExecTask() {
	task, err := GetTask()
	if err != nil {
		a_logger.Error(err.Error())
		return
	}
	if strings.ToLower(task.Type) == "get" {
		qStr := QueryString(task.Field, "get")
		if task.Method[0] == '/' || len(task.Method) == 0 {
			qStr = task.Ip + ":" + task.Port + task.Method + "?" + qStr
		} else {
			qStr = task.Ip + ":" + task.Port + "/" + task.Method + "?" + qStr
		}
		fmt.Println(qStr)
	}
	if strings.ToLower(task.Type) == "post" {
		var qApi = ""
		if task.Method[0] == '/' || len(task.Method) == 0 {
			qApi = task.Ip + ":" + task.Port + task.Method
		} else {
			qApi = task.Ip + ":" + task.Port + "/" + task.Method
		}
		qStr := QueryString(task.Field, "post")
		fmt.Println(qStr)
		fmt.Println(qApi)
	}
}
