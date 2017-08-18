/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Types Drivered by Golang
 * 2017/8/2, by Tan Haipeng, create
 */

package main

type RetType struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type StatType struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data SysType `json:"data"`
}

type SysType struct {
	Stat string `json:"stat"`
	Time string `json:"time"`
}

type TaskType struct {
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Method  string `json:"method"`
	Count   int `json:"count"`
	Timeout int `json:"timeout"`
	Type    string `json:"type"`
	Field   []FieldType `json:"data"`
}

type FieldType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
