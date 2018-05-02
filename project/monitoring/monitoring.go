package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"time"
	"net/url"
	"fmt"
)

type ModuleStatus struct {
	description string
	status      string
}

func curl(moduleName string,environment string) {
	slackUrl := "https://hooks.slack.com/services/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"
	//初始化http.Client对象
	client := &http.Client{}
	//post请求
	postValues := url.Values{}
	message := fmt.Sprintf("environment[%v]  module[%v] status is error",environment,moduleName)
	postValues.Add("payload", "{\"text\": \""+message+"\", \"icon_emoji\": \":ghost:\"}")
	resp, err := client.PostForm(slackUrl, postValues)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}
}

func check(moduleName string,environment string,url string) bool{
	//log.Printf("check environment[%v]  module[%v] url[%v]\n",environment,moduleName,url)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	//log.Println("Get")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	//log.Println("ReadAll")
	var data map[string]interface{} // TopTracks
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		return false
	}
	//log.Println("Unmarshal")
	if data["status"] == "UP" {
		log.Printf("environment[%v]  module[%v] status is up \n",environment,moduleName)
		return true
	}
	log.Printf("environment[%v]  module[%v] status is error \n",environment,moduleName)
	log.Println("error",err)
	return false
}

func checkEnvironment(moduleArray [] string) {
	for _, module := range moduleArray {
		log.Printf("check module[%v] start",module)
		//url := "http://cloud-dt.deja.fashion/style-tinder/health"
		dtUrl := "http://cloud-dt.deja.fashion/" + module + "/health"
		if !check("dt",module,dtUrl){
			curl(module,"dt")
		}
		dpUrl := "http://cloud-dp.deja.fashion/" + module + "/health"
		if !check("dp",module,dpUrl){
			curl(module,"dp")
		}
		productionUrl := "https://cloud.deja.fashion/" + module + "/health"
		if !check("production",module,productionUrl){
			curl(module,"production")
		}
		log.Printf("check module[%v] end",module)
	}

}

func main() {
	log.Println("monitoring begin")
	//初始化定时器
	//ticker := time.NewTicker(300 * time.Second)
	ticker := time.NewTicker(5 * time.Second)
	//moduleArray := []string{"app-config","auth"}
	moduleArray := []string{"app-config","auth","cashback","customer","favourite","id-generator","inventory","invoice","legacy-db","message","ocb-syncer","ocr","order","payment","scheduler","shop","shopping-bag","stripe","style-tinder","wardrobe"}
	log.Printf("initial moduleArray : %v \n",moduleArray)
	var liveStatus int = 0;
	for _ = range ticker.C {
		log.Println("I'm alive.")
		liveStatus++
		if(liveStatus == 60){
			liveStatus = 0
			checkEnvironment(moduleArray)
		}

	}
	log.Println("monitoring end")
	//退出程序
	os.Exit(0)
}