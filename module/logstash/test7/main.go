package main

import (
	"sillyhat/module/logstash/test7/test7"
)

func main() {
	log := test7.NewLogger()
	log.Info("test")
}