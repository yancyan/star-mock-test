package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {

}

func MockServer(address string) {

	//simpleServer()

	router := InitRouter()
	er := router.Run(address)
	if er != nil {
		log.Fatal("server error...", er)
	}

	//server := initServer(address, router)
	//err := server.ListenAndServe(); if err != nil {
	//	requests.Logger.Fatal("server listen error, ", err)
	//}
}

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func simpleServer() {
	r := gin.Default()
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("logs/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 心跳
	r.GET("/status", func(cxt *gin.Context) {
		cxt.String(http.StatusOK, "status %s", "ok...")
	})

	//获取Get参数
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		age := c.Query("age")
		fmt.Printf("name: %s, age: %s \n", name, age)

		c.String(http.StatusOK, "Hello %s %s", name, age)
	})

	//获取Post参数
	r.POST("/form", func(c *gin.Context) {
		msg := c.PostForm("message")
		name := c.DefaultPostForm("name", "Guest")

		c.JSON(http.StatusOK, gin.H{
			"message": msg,
			"name":    name,
		})

	})

	er := r.Run(":9977")
	if er != nil {
		log.Fatal("server error...", er)
	}
}
