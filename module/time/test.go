package main

import (
	"log"
	"time"
)

func getTimeInMillis() string{
	return ""
}

func main() {
	//log.Println(time.Now().Unix())
	//log.Println(time.Now().Unix()/60)

	//time.Now().Format("2006-01-02 15:04:05")

	startTime := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2018, 1, 15, 0, 0, 0, 0, time.Local)
	test1 := endTime.Unix() - startTime.Unix()
	test2 := endTime.Unix() - startTime.Unix()
	log.Println(test1)
	log.Println(test2)
	log.Println(time.Now().Unix())
	log.Println(time.Now().UnixNano()/1e6)
	//log.Println(getTimeInMillis())

	log.Println(time.Now().Unix()) //获取当前秒
	log.Println(time.Now().UnixNano())//获取当前纳秒
	log.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒
	log.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒
	c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型
	log.Println(c.String()) //输出当前英文时间戳格式
}
