package main

import (
	"time"
	"log"
)

func test1Func(timeout time.Duration){
	var after <-chan time.Time
	after = time.After(timeout)
	i := 0
	for{
		log.Println("等待a中的数据，10秒后没有数据则超时")
		i++
		select {
		case i:
			log.Println(i)
		case <-after:
			log.Println("timeout.")
			return
		}
	}
}

func test1(){
	test1Func(10 * time.Second)
	log.Println("test1Func end.")
}

func main() {
	test1()
}
