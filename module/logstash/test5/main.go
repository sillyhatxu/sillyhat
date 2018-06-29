package main

import (
	log "sillyhat/module/logstash/test5/testpublic"
)

func init() {
	//log.Initial("Golang Test Module")
	log.Initial("Golang Test Module")
}

func testLog()  {
	log.Info("hahaha")
}
func main() {
	go testLog()

	log.Print("test Print")
	log.Info("test Info")
	log.Debug("test Debug")
	log.Warn("test Warn")
	log.Warning("test Warning")
	log.Error("test Error")
}
