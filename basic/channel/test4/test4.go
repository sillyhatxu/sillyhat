package main

import "time"
import "fmt"

var index int
func main() {
	index = 0
	c1 := make(chan string, 1)

	go func() {
		if index == 10{
			time.Sleep(time.Second * 2)
		}
		index++
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}
