package main

import "log"

type Callback func(x,y int) int

type ClientMethod struct{
	method string
	callback Callback
}

func test(x,y int,callback Callback) int {
	return callback(x,y)
}
//method    msg   callback


func main() {
	log.Fatal("consul client error : ", 123)
	log.Println("534363")

}



func testCallback1(msg string){
	log.Println(msg)
}

func testCallback2(msg string){
	log.Println(msg)
}