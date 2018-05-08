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
	"github.com/golang/glog"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
)

var (
	enviroment string
	configDir string
    configMap = make(map[string]interface{})
)

type CommandDTO struct {

	CommandName string `json:"commandName"`

	CommandBody string `json:"commandBody"`
}

type CommandBodyDTO struct {

	Slack *CommandBodySlack `json:"slack"`
}

type CommandBodySlack struct {

	Pretext string `json:"pretext"`

	MessageUrl string `json:"messageUrl"`

	Time int64 `json:"time"`

	Color string `json:"color"`

	SlackDetailArray []CommandBodySlackDetail `json:"slackDetailList"`
}

type CommandBodySlackDetail struct {

	Title string `json:"title"`

	Value string `json:"value"`
}
//{"attachments":[{"color":"color","pretext":"pretext","ts":1525747854167,"fields":[{"title":"title1","value":"value1","short":false},{"title":"title2","value":"value2","short":false}]}]}
type Attachments struct {
	SlackArray [] Slack `json:"attachments"`
}

type Slack struct {
	Color string `json:"color"`

	Pretext string `json:"pretext"`

	Time int64 `json:"ts"`

	FieldArray []SlackDetail `json:"fields"`
}

type SlackDetail struct {
	Title string `json:"title"`

	Value string `json:"value"`

	Short bool `json:"short"`
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
	configDir = dir + "/project/message/" + getConfigName()
	//if enviroment == "dev"{
	//	dir, err := os.Getwd()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	configDir = dir + "/project/message/" + getConfigName()
	//}else{
	//	configDir = "/go/" + getConfigName()
	//}
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

func getCommandBodySlack(commandBodySlack *CommandBodySlack) *CommandBodySlack {
	if(commandBodySlack.Time == 0){
		commandBodySlack.Time = time.Now().Unix()
	}
	if(commandBodySlack.Color == ""){
		commandBodySlack.Color = configMap["message.default.color"].(string)
	}
	if commandBodySlack.MessageUrl == ""{
		commandBodySlack.MessageUrl = configMap["message.default.messageurl"].(string)
	}
	return commandBodySlack
}
//{"attachments":[{"color":"color","pretext":"pretext","ts":1525701011018,"fields":[{"title":"title1","value":"value1","short":false},{"title":"title2","value":"value2","short":false}]}]}
//{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"aaa\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1525697855,\"color\":\"#36a64f\"}}"}
func concumer(messageJson string){
	//{"attachments":[{"color":"color","pretext":"pretext","ts":1525747854167,"fields":[{"title":"title1","value":"value1","short":false},{"title":"title2","value":"value2","short":false}]}]}
	var command CommandDTO
	json.Unmarshal([]byte(messageJson), &command)
	log.Println(command)

	var commandBody CommandBodyDTO
	json.Unmarshal([]byte(command.CommandBody), &commandBody)

	//attachments := &Attachments{{Pretext:commandBody.Slack.Pretext,Color:commandBody.Slack.Color}}
	log.Println(commandBody)
	var fieldArray []SlackDetail
	for i:= 0;i < len(commandBody.Slack.SlackDetailArray);i++{
		commandBodySlackDetail := commandBody.Slack.SlackDetailArray[i]
		fieldArray = append(fieldArray,SlackDetail{Title:commandBodySlackDetail.Title,Value:commandBodySlackDetail.Value,Short:true})
	}
	commandBodySlack := getCommandBodySlack(commandBody.Slack)
	attachments := &Attachments{
		SlackArray:[]Slack{
			{
				Pretext: commandBodySlack.Pretext,
				Color: commandBodySlack.Color,
				Time: commandBodySlack.Time,
				FieldArray:fieldArray,
			},
		},
	}
	attachmentsJson, err := json.Marshal(attachments)
	check(err)
	log.Println(string(attachmentsJson))
	curl(commandBodySlack.MessageUrl,string(attachmentsJson))
}

func startKafka()  {
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
		glog.Errorf("Failed open consumer: %v", err)
		return
	}
	defer c.Close()
	go func() {
		for err := range c.Errors() {
			glog.Errorf("Error: %s\n", err.Error())
		}
	}()

	go func() {
		for note := range c.Notifications() {
			glog.Infof("Rebalanced: %+v\n", note)
		}
	}()

	for msg := range c.Messages() {
		log.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		log.Println()
		concumer(string(msg.Value))
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}


func main() {
	flag.StringVar(&enviroment, "enviroment", "dev", "enviroment")
	flag.Parse()
	log.Printf("enviroment : %v \n",enviroment)
	readConfig()
	startHealthApi()
	startKafka()
	//fmt.Printf("%d \n", time.Now().UnixNano() / int64(time.Millisecond))
	//fmt.Printf("%d \n", time.Now().UnixNano() / int64(time.Microsecond))
	//fmt.Println(time.Now().Unix())//获取时间戳
	//testJson := `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"Payment success.\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1525702528,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"User\",\"value\":\"vixathep phoyphailin\"},{\"title\":\"Order Number\",\"value\":\"225428419044631\"},{\"title\":\"Brand\",\"value\":\"Zara\"},{\"title\":\"Items\",\"value\":\"2\"},{\"title\":\"Total Amount\",\"value\":\"17.06\"}]}}"}`
	//concumer(testJson)
	//testJson2 := `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"Time null.\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"User\",\"value\":\"vixathep phoyphailin\"},{\"title\":\"Order Number\",\"value\":\"225428419044631\"},{\"title\":\"Brand\",\"value\":\"Zara\"},{\"title\":\"Items\",\"value\":\"2\"},{\"title\":\"Total Amount\",\"value\":\"17.06\"}]}}"}`
	//concumer(testJson2)
	//testJson3 := `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"url null.\",\"time\":1525702528,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"User\",\"value\":\"vixathep phoyphailin\"},{\"title\":\"Order Number\",\"value\":\"225428419044631\"},{\"title\":\"Brand\",\"value\":\"Zara\"},{\"title\":\"Items\",\"value\":\"2\"},{\"title\":\"Total Amount\",\"value\":\"17.06\"}]}}"}`
	//concumer(testJson3)
	//testJson4 := `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"color null.\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1525702528,\"slackDetailList\":[{\"title\":\"User\",\"value\":\"vixathep phoyphailin\"},{\"title\":\"Order Number\",\"value\":\"225428419044631\"},{\"title\":\"Brand\",\"value\":\"Zara\"},{\"title\":\"Items\",\"value\":\"2\"},{\"title\":\"Total Amount\",\"value\":\"17.06\"}]}}"}`
	//concumer(testJson4)
	//testJson5 := `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"fieldList null.\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1525702528,\"color\":\"#36a64f\"}}"}`
	//concumer(testJson5)
}


type ProjectStatus struct {
	Description string   `json:"description"`
	Status  string   `json:"status"`
}

func health(response http.ResponseWriter, request *http.Request) {
	projectStatus := ProjectStatus{Description:"Golang project client status",Status:"UP"}
	log.Printf("Description : %v; Status : %v\n",projectStatus.Description,projectStatus.Status)
	json.NewEncoder(response).Encode(projectStatus)
}

func startHealthApi()  {
	log.Println("start health api")
	router := mux.NewRouter()
	router.HandleFunc("/health", health).Methods("GET")
	log.Fatal(http.ListenAndServe(":18001", router))
}

//file, err := os.Open(".")
//if err != nil {
//	log.Fatalf("failed opening directory: %s", err)
//}
//defer file.Close()
//
//list,_ := file.Readdirnames(0) // 0 to read all files and folders
//for _, name := range list {
//	fmt.Println(name)
//}
//fmt.Println("\n------------------\n")
