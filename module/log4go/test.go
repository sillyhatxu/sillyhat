package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(10000)  // 创建一个日志记录器，参数为缓冲区的大小
	log.SetLogger("console", "")  // 设置日志记录方式：控制台记录
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	log.Info("this %s cat is %v years old", "yellow", 3)
	//an official log.Logger
	//log := logs.GetLogger()
	//logs.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ocb-syncer").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")

	logs.GetLogger("ocb-syncer").Println("Send Legacy DB event : success")



	//2018-05-16 12:14:57.802 - INFO [app-config][io-12000-exec-8] f.d.a.controller.AppConfigController     [   217]: Send Legacy DB event : success
}