package main


import (
	"os"
	"log"
	"io"
	"bufio"
	"strings"
	"github.com/bsm/sarama-cluster"
	"time"
	"github.com/Shopify/sarama"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"flag"
)

var (
	enviroment string
	configDir string
	configMap = make(map[string]interface{})
)

type EventDTO struct {

	EventName string `json:"eventName"`

	EventBody string `json:"eventBody"`
}

type EventBodyDTO struct {

	Uid int64 `json:"uid"`

	PhoneNumber string `json:"phoneNumber"`

	OperatorTime time.Time `json:"operatorTime"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getConfigName() string {
	return "config-"+enviroment+".ini"
}

func getConfigDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//configDir = dir + "/" + getConfigName()
	//configDir = dir + "/project/message/" + getConfigName()
	if enviroment == "dev"{
		configDir = dir + "/module/kafkatest/" + getConfigName()
	}else{
		configDir = dir + "/" + getConfigName()
	}
	log.Printf("configDir : [%v]\n",configDir)
	return configDir
}

func readConfig() {
	//打开这个ini文件
	f, _ := os.Open(getConfigDir())
	//读取文件到buffer里边
	buf := bufio.NewReader(f)
	for {
		//按照换行读取每一行
		l, err := buf.ReadString('\n')
		//相当于PHP的trim
		line := strings.TrimSpace(l)
		//判断退出循环
		if err != nil {
			if err != io.EOF {
				//return err
				panic(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
			//匹配[db]然后存储
		case line[0] == '[' && line[len(line)-1] == ']':
			section := strings.TrimSpace(line[1 : len(line)-1])
			log.Println(section)
		default:
			//dnusername = xiaowei 这种的可以匹配存储
			i := strings.IndexAny(line, "=")
			configMap[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])
		}
	}
	//循环输出结果
	log.Println("--------------- config start ---------------")
	for k, v := range configMap {
		log.Println(k,":", v)
	}
	log.Println("--------------- config end ---------------")
}

func curl(slackURL string,slackMessage string) {
	//初始化http.Client对象
	client := &http.Client{}
	//post请求
	postValues := url.Values{}
	log.Printf("url[%v]  slackMessage[%v]\n",slackURL,slackMessage)
	postValues.Add("payload", slackMessage)

	//Content-type: application/json
	resp, err := client.PostForm(slackURL, postValues)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}
}
//{"attachments":[{"color":"color","pretext":"pretext","ts":1525701011018,"fields":[{"title":"title1","value":"value1","short":false},{"title":"title2","value":"value2","short":false}]}]}
//{"eventName":"sendSlack","eventBody":"{\"slack\":{\"pretext\":\"aaa\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1525697855,\"color\":\"#36a64f\"}}"}
func concumer(messageJson string){
	log.Println(messageJson)
	//{"attachments":[{"color":"color","pretext":"pretext","ts":1525747854167,"fields":[{"title":"title1","value":"value1","short":false},{"title":"title2","value":"value2","short":false}]}]}
	//var event EventDTO
	//json.Unmarshal([]byte(messageJson), &event)
	//log.Println(event)
	//
	//var eventBody EventBodyDTO
	//json.Unmarshal([]byte(event.EventBody), &eventBody)
	//
	////attachments := &Attachments{{Pretext:eventBody.Slack.Pretext,Color:eventBody.Slack.Color}}
	//log.Println(eventBody)

}

func startKafka()  {
	log.Println("start kafka service")
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	//config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	groupID:= configMap["kafka.group"].(string)
	nodeList:= configMap["kafka.node.list"].(string)
	topicList:= configMap["kafka.topic.list"].(string)

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
		log.Println()
		//concumer(string(msg.Value))
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

//-enviroment=dt
func main() {
	flag.StringVar(&enviroment, "enviroment", "dev", "enviroment")
	flag.Parse()
	log.Printf("enviroment : %v \n",enviroment)
	readConfig()
	startKafka()
}


//2018/05/17 18:07:19 Rebalanced: &{rebalance start map[] map[] map[]}
//2018/05/17 18:07:19 Rebalanced: &{rebalance OK map[LoginEvent:[0 1]] map[] map[LoginEvent:[0 1]]}
//2018/05/17 18:07:19 Partition:0, Offset:42, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":78903,\"phoneNumber\":\"+6597729756\",\"operatorTime\":1526538097876}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:0, Offset:43, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84344,\"phoneNumber\":\"+6591500498\",\"operatorTime\":1526546112845}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:0, Offset:44, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84349,\"phoneNumber\":\"+6592717588\",\"operatorTime\":1526549573509}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:0, Offset:45, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84262,\"phoneNumber\":\"+6583324064\",\"operatorTime\":1526550631869}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:0, Offset:46, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":78903,\"phoneNumber\":\"+6597729756\",\"operatorTime\":1526551316587}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:41, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84243,\"phoneNumber\":\"+6592717588\",\"operatorTime\":1526537332107}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:42, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84330,\"phoneNumber\":\"+6591500498\",\"operatorTime\":1526541494787}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:43, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84344,\"phoneNumber\":\"+6591500498\",\"operatorTime\":1526549262384}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:44, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":84349,\"phoneNumber\":\"+6592717588\",\"operatorTime\":1526550166731}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:45, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":83038,\"phoneNumber\":\"+6596125706\",\"operatorTime\":1526551205124}"}
//2018/05/17 18:07:19
//2018/05/17 18:07:19 Partition:1, Offset:46, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":78903,\"phoneNumber\":\"+6597729756\",\"operatorTime\":1526551510608}"}
//2018/05/17 18:07:19
//2018/05/17 18:12:23 Partition:0, Offset:47, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":78903,\"phoneNumber\":\"+6597729756\",\"operatorTime\":1526551881636}"}
//2018/05/17 18:12:23
//2018/05/17 18:14:18 Partition:1, Offset:47, Key:, Value:{"eventName":"login","eventBody":"{\"uid\":78903,\"phoneNumber\":\"+6597729756\",\"operatorTime\":1526551996389}"}
//2018/05/17 18:14:18