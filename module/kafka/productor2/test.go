package main

import (
	"github.com/Shopify/sarama"
	"time"
	"strings"
	"github.com/golang/glog"
	"strconv"
	"log"
	"os"
	"fmt"
)

var (
	topics = "MessageCommandLocal"
	kafkaProducer sarama.AsyncProducer
)

func getValue(threadName string,i int) string {
	return `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"` + threadName+`---`+strconv.Itoa(i) + `\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1526630475,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"title1\",\"value\":\"value1\"},{\"title\":\"title2\",\"value\":\"value2\"},{\"title\":\"title3\",\"value\":\"value3\"},{\"title\":\"title4\",\"value\":\"value4\"}]}}"}`
}


func initial()  {
	log.Println("initial")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	//test, err := sarama.NewAsyncProducer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), config)
	kafkaProducer,_ = sarama.NewAsyncProducer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), config)
	//defer kafkaProducer.Close()
	//if err != nil {
	//	return
	//}
	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					glog.Errorln(err)
				}
			case <-success:
			}
		}
	}(kafkaProducer)
}

func sendEvent(value string){
	fmt.Fprintln(os.Stdout, value)
	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(value),
	}
	kafkaProducer.Input() <- msg
}

func testProducer(threadName string,min int,max int) {
	for i := min;i <= max ;i++  {
		sendEvent(getValue(threadName,i))
	}
}

func main() {
	initial()
	//sendEvent(getValue("test",1))
	log.Println("sendEvent")
	testProducer("AAAAAAAAAA",1001,8000)
	//go testProducer("BBBBBBBBBB",2001,3000)
	//go testProducer("CCCCCCCCCC",3001,4000)
	//go testProducer("DDDDDDDDDD",4001,5000)
	//go testProducer("EEEEEEEEEE",5001,6000)
	//time.Sleep(50000000000)
}
