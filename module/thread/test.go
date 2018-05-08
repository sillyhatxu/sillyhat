package main

import (
	"fmt"
	"time"
)

func loop(threadName string) {
	fmt.Printf("%v \n",threadName)
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d ", i)
	}
}


func main() {
	go loop("A")
	loop("B")
	time.Sleep(time.Second*10) // 停顿一秒
}