package test1

import (
	"fmt"
	"time"
)

func main(){
	//初始化定时器
	ticker := time.NewTicker(300 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
	}
}