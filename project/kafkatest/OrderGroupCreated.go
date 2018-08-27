package main


import (
	log "sillyhat-golang-tool/sillyhat_log/logrus"
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	moduleName = "message"
)

func init(){
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetLevel(log.InfoLevel)
	log.SetModuleName(moduleName)
	log.SetHookType(log.Elasticsearch)
	log.Printf("-------------------- initial server enviroment %v--------------------\n",moduleName)
	/***** system enviroment config *****/
}

func main() {
	log.Printf("-------------------- start server %v--------------------\n",moduleName)
	router := gin.Default()
	/***** system initial end *****/
	go startKafka()
	/***** web api start *****/
	router.Run(":8800")
	/***** web api end *****/
}

func startKafka()  {
	log.Info("start kafka service")
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	//enviroment.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	groupID := "test"
	nodeList := "172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090"
	topicList := "OrderGroupEvent"

	c, err := cluster.NewConsumer(strings.Split(nodeList, ","), groupID, strings.Split(topicList, ","), config)
	if err != nil {
		log.Printf("Failed open consumer: %v", err)
		return
	}
	defer c.Close()
	go func() {
		for err := range c.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	go func() {
		for note := range c.Notifications() {
			log.Printf("Rebalanced: %v \n", note)
		}
	}()

	for msg := range c.Messages() {
		log.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		//concumer(msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}

}
