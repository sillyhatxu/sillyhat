package main


import (
	"os"
	"io"
	"bufio"
	"strings"
	"time"
	"encoding/json"
	"database/sql"
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"flag"
)

var (
	db = &sql.DB{}
	enviroment string
	configDir string
	configMap = make(map[string]interface{})
)

type EventDTO struct {

	EventName string `json:"eventName"`

	EventBody string `json:"eventBody"`
}

type EventBodyDTO struct {

	CustomerId int64 `json:"customerId"`

	UpdateShopItemPrice *UpdateShopItemPriceMessageDTO `json:"updateShopItemPrice"`
}

type UpdateShopItemPriceMessageDTO struct {

	Id int64 `json:"shop_item_id"`

	CurrentPrice int64 `json:"current_price"`

	OriginalPrice int64 `json:"original_price"`

	Currency string `json:"currency"`
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
		log.Error(err)
	}
	if enviroment == "dev"{
		configDir = dir + "/project/modifyprice/" + getConfigName()
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

func concumer(messageJson string){
	var event EventDTO
	json.Unmarshal([]byte(messageJson), &event)
	log.Println(event)
	if(event.EventName == "updateShopItemPrice"){
		var eventBody EventBodyDTO
		json.Unmarshal([]byte(event.EventBody), &eventBody)

		log.Info(eventBody)
		concumerBody(eventBody)
	}
}

func concumerBody(eventBody EventBodyDTO){
	tx,_ := db.Begin()
	defer tx.Commit()
	row := tx.QueryRow("SELECT COUNT(1) as count FROM ocb_syncer.shop_item_stop_refresh WHERE is_deleted = false AND shop_item_id = ?",eventBody.UpdateShopItemPrice.Id)
	var count int
	row.Scan(&count)
	log.Printf("result count : %d",count)
	if count == 0 {
		currency := "S$"
		if eventBody.UpdateShopItemPrice.Currency != "SGD" {
			currency = eventBody.UpdateShopItemPrice.Currency
		}
		stm,err := db.Prepare("UPDATE ocb_syncer.shop_item_temp SET current_price = ?,original_price = ?,currency = ?,status = ?,validate_status = ?,last_modified_by = ?,last_modified_date = sysdate() WHERE id = ?")
		if err != nil{
			log.Printf("Update shop item price error: %s\n", err.Error())
		}else{
			stm.Exec(eventBody.UpdateShopItemPrice.CurrentPrice,eventBody.UpdateShopItemPrice.OriginalPrice,currency,1000,0,eventBody.CustomerId,eventBody.UpdateShopItemPrice.Id)
			stm.Close()
			log.Printf("UPDATE ocb_syncer.shop_item_temp SET current_price = %v,original_price = %v,currency = %v,status = %v,validate_status = %v,last_modified_by = %v,last_modified_date = sysdate() WHERE id = %v",eventBody.UpdateShopItemPrice.CurrentPrice,eventBody.UpdateShopItemPrice.OriginalPrice,currency,1000,0,eventBody.CustomerId,eventBody.UpdateShopItemPrice.Id)
		}

	}else {
		log.Printf("stop refresh shop item price %d\n",eventBody.UpdateShopItemPrice.Id)
	}

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
		concumer(string(msg.Value))
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

func init(){
	flag.StringVar(&enviroment, "enviroment", "dev", "enviroment")
	flag.Parse()
	log.Printf("enviroment : %v \n",enviroment)
	readConfig()
	db,_ = sql.Open("mysql", configMap["database.dataSourceName"].(string))
	connMaxLifetime := 10
	maxIdleConns := 10
	maxOpenConns := 100
	log.Printf("connMaxLifetime:%d\n", connMaxLifetime)
	log.Printf("maxIdleConns:%d\n", maxIdleConns)
	log.Printf("maxOpenConns:%d\n", maxOpenConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
}

//-enviroment=dt
func main() {
	startKafka()
	//testJson := `{"eventName":"updateShopItemPrice","eventBody":"{\"customerId\":1,\"updateShopItemPrice\":{\"shop_item_id\": 5426900,\"current_price\": 2000,\"original_price\": 3990,\"currency\": \"SGD\"}}"}`
	//concumer(testJson)
}