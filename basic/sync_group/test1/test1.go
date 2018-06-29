package main

import (
	"sync"
	"strconv"
	"time"
	"fmt"
)

func testPring(threadName string)  {
	for i := 1; i <= 20;i++  {
		fmt.Println(threadName + " [" + strconv.Itoa(i) + "]")
		time.Sleep(1*time.Second)
	}
	fmt.Println(threadName + " finish.")
}


func main() {
	var wg sync.WaitGroup
	for i:=1;i <= 10;i++{
		threadName := "Thread " + strconv.Itoa(i)
		wg.Add(1)
		go func(threadName string) {
			testPring(threadName)
			defer wg.Done()
		}(threadName)
	}
	wg.Wait()
	fmt.Println("finish")
	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//<-done
}
