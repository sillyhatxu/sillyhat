package main


import (
	"github.com/Shopify/sarama"
	"time"
	"strings"
	"github.com/golang/glog"
	"os"
	"fmt"
	"strconv"
)

var (
	topics = "MessageCommandLocal"
)

//{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"ssssssss\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1526630475,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"title1\",\"value\":\"value1\"},{\"title\":\"title2\",\"value\":\"value2\"},{\"title\":\"title3\",\"value\":\"value3\"},{\"title\":\"title4\",\"value\":\"value4\"}]}}"}
//type SendSlackCommand struct {
//
//}

func getValue(i int) string {
	return `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"` + strconv.Itoa(i) + `\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1526630475,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"title1\",\"value\":\"value1\"},{\"title\":\"title2\",\"value\":\"value2\"},{\"title\":\"title3\",\"value\":\"value3\"},{\"title\":\"title4\",\"value\":\"value4\"}]}}"}`
}

func asyncProducer(i int) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), config)
	defer p.Close()
	if err != nil {
		return
	}

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
	}(p)

	value := getValue(i)
	fmt.Fprintln(os.Stdout, value)
	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(value),
	}
	p.Input() <- msg
}
//{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"ssssssss\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1526630475,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"title1\",\"value\":\"value1\"},{\"title\":\"title2\",\"value\":\"value2\"},{\"title\":\"title3\",\"value\":\"value3\"},{\"title\":\"title4\",\"value\":\"value4\"}]}}"}

func producer(min int,max int) {
	for i := min;i <= max ;i++  {
		asyncProducer(i)
	}
}

func main() {
	go producer(1001,2000)
	go producer(2001,3000)
	go producer(3001,4000)
	go producer(4001,5000)
	go producer(5001,6000)
	time.Sleep(50000000000)
}