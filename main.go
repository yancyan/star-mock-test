package main

import "kingdee-test/requests"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	requests.TestResourceForm()
	requests.TestC021GetStock()

	//for i := 0; i < 100; i++ {
	//	//go requests.Login()
	//	go requests.TestC021GetStock()
	//}
	//time.Sleep(1000*time.Second)

	// 启动mock服务
	//router.MockServer(":7788")
}
