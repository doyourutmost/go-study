package main

import (
	"learn-go-pro/log_test"
)

var log log_test.Logger // 声明一个全局的接口变量

// 测试我们自己写 的日志库
func main() {
	log = log_test.NewConsoleLogger("Info")                                    // 终端日志实例
	log = log_test.NewFileLogger("Info", "./", "zhoulinwan.log", 10*1024*1024) // 文件日志实例

	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(2 * time.Second)
	}
}
