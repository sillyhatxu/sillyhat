package main

import (
	"io/ioutil"
	"net/http"
	"time"
	"log"
)


func httpGet1(uid string) {
	resp, err := http.Get("http://cloud.deja.fashion/shop/recommend")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	log.Println(string(body))
}

func httpGet2(uid string)  {
	//timeout := time.Duration(50 * time.Millisecond)//超时时间50ms
	timeout := time.Duration(30 * time.Second)//超时时间50ms
	client := &http.Client{Timeout: timeout}
	//生成要访问的url
	url := "http://cloud.deja.fashion/shop/recommend"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("uid", uid)
	//reqest.Header.Add("User-Agent", "xxx")
	//reqest.Header.Add("X-Requested-With", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		// handle error
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}
	log.Println(string(body))
	defer response.Body.Close()
}

func timeElapsed(uid string){
	t1 := time.Now() // get current time
	httpGet2(uid)
	elapsed := time.Since(t1)
	log.Println("App elapsed: ", elapsed)
}


func main() {
	var idArray []string
	//idArray = append(idArray, "78255","63200","46127","68133","24276","305268","293547","20271","245327","260086")
	idArray = append(idArray, "245327","260086")
	for{
		log.Println("-------------------- test start --------------------")
		for _, v := range idArray {
			go timeElapsed(v)
		}
		time.Sleep(time.Second*10)
		log.Println("-------------------- test end --------------------")
	}
	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//<-done
}
