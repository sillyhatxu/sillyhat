package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(5*time.Second)
	c <- sum // send sum to c
}


func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	fmt.Println("sum start")
	x, y := <-c, <-c // receive from c
	fmt.Println("sum end")
	fmt.Println(x, y, x+y)

}