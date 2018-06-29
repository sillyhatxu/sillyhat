package other

import (
	log "sillyhat/module/logstash/test6/logger"
)



func TestLogger()  {
	log.Log().Print("other test Print")
	log.Log().Info("other test Info")
	log.Log().Debug("other test Debug")
	log.Log().Warn("other test Warn")
	log.Log().Warning("other test Warning")
	log.Log().Error("other test Error")
}
