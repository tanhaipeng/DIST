/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Control Module Drivered by Golang
 * 2017/7/25, by Tan Haipeng, create
 */

package main

import (
	"fmt"
	"net/http"
)

func init() {
	initConfig()
	initLogger()
	initRouter()
}

func initRouter() {
	http.HandleFunc("/stat", getSlaveStat)
}

func main() {
	fmt.Println("DIST Control Module")
	http.ListenAndServe(":8010", nil)
}
