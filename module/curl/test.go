package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
)

func main() {
	slackUrl := "https://hooks.slack.com/services/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"
	//初始化http.Client对象
	client := &http.Client{}

	//post请求
	postValues := url.Values{}
	postValues.Add("payload", "{\"text\": \"This is posted to #module_notifaction and comes from a bot named webhookbot.\", \"icon_emoji\": \":ghost:\"}")
	resp, err := client.PostForm(slackUrl, postValues)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}

	//post请求
	//postValues := url.Values{}
//	postValues.Add("publicKey", "")
//	postValues.Add("privateKey", `----nGDd4/mujoJBr5mkrw
//DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
//AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
//-----END RSA PRIVATE KEY-----`)
//	postValues.Add("info", "")
//	postValues.Add("message", `DkCxcs0z6Z03uHWOHOASf2xen+7oNoSad+KG2ss0hkE79211GlgjepmMFRW4zLiF51pVYHHOBFDYYJrnokq5d0ceKYY6ONzbBYKCJMzD7guN3qMYf48Cl9g0bDVb1oMbuN2PstzORe800Q72moQaHVRPiqh7VZ6NCXnkLrtnY64=`)



}



//curl -X POST --data-urlencode "payload={\"channel\": \"#module_notifaction\", \"username\": \"webhookbot\", \"text\": \"This is posted to #module_notifaction and comes from a bot named webhookbot.\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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
//curl -X POST --data-urlencode "payload={\"text\": \"This is posted to #module_notifaction and comes from a bot named webhookbot.\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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
