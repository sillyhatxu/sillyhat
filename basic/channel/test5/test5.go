package main

import (
	"time"
	"fmt"
)

func TestChan() {
	c := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("channel")
		c <- 48
	}()

	fmt.Println("sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("start")
	fmt.Println(<- c)
	// 保持持续运行
	holdRun()
}

func holdRun() {
	time.Sleep(1 * time.Hour)
}

func main() {
	TestChan()
}
