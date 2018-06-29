package main

import (
	log "sillyhat/module/logstash/test6/logger"
	"github.com/pkg/errors"
	"sillyhat/module/logstash/test6/other"
)

func init() {
	log.Initial("Test Module")
}


func testLog()  {
	log.Log().Info("hahaha")
}

func testErrorLog()  {
	log.Log().Error("error",errors.New("a error message"))
}

func main() {
	go testLog()
	go testErrorLog()
	other.TestLogger()
	log.Log().Print("test Print")
	log.Log().Info("test Info")
	log.Log().Debug("test Debug")
	log.Log().Warn("test Warn")
	log.Log().Warning("test Warning")
	log.Log().Error("test Error")
}