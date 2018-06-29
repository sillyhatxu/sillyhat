package main

import (
	"fmt"
	"log"

	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

func consulCheck2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

func registerServer2() {

	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	checkPort := 18080

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "SillyHat_Id"
	registration.Name = "SillyHat"
	registration.Port = 19527
	registration.Tags = []string{"SillyHatTags"}
	registration.Address = "127.0.0.1"
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	http.HandleFunc("/check", consulCheck2)
	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)

}

func main() {
	registerServer2()
}