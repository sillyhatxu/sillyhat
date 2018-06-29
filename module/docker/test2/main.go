package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"sillyhat-golang-tool/sillyhat_scheduler"
)

func test1()  {
	log.Println("test1")
}

func test2()  {
	log.Println("test2")
}

func test3(module string)  {
	log.Printf("test3 %v\n",module)
}

func main() {
	sillyhat_scheduler.InitialScheduler(test1,"00:00:00","5s")
	sillyhat_scheduler.InitialScheduler(test2,"00:00:00","10s")
	sillyhat_scheduler.InitialScheduler(test3,"00:00:00","15s","test scheduler params")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}