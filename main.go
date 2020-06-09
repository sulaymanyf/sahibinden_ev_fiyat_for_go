package main

import (
	"./drivers"
	"./logic"
	"./server"
	"time"

	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	// 服务停止时清理数据库链接
	defer drivers.MysqlDb.Close()
	startTimer(logic.SatilikJob)
	startTimer(logic.KiralikJob)
	// 启动服务
	server.Run(HttpServer)

}

func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
