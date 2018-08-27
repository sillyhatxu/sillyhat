package main

import (
	"fmt"
	"log"
	"net"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

const RECV_BUF_LEN = 1024

func main() {

	client, err := consulapi.NewClient(consulapi.DefaultConfig())

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	for {

		time.Sleep(time.Second * 3)
		var services map[string]*consulapi.AgentService
		var err error

		services, err = client.Agent().Services()

		if nil != err {
			log.Println("in consual list Services:", err)
			continue
		}

		if _, found := services["serverNode_1"]; !found {
			log.Println("serverNode_1 not found")
			continue
		}

		sendData(services["serverNode_1"])
		break
	}
}

func sendData(service *consulapi.AgentService) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN)
	log.Println("-------------------------------")
	i := 0
	//for {
		i++
		log.Println("------------2-------------------")
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))//调用server
		if err != nil {
			println("Write Buffer Error:", err.Error())
			//break
		}

		n, err = conn.Read(buf)//读取调用server返回信息
		if err != nil {
			println("Read Buffer Error:", err.Error())
			//break
		}
		log.Println("get:", string(buf[0:n]))

		//等一秒钟
		time.Sleep(10 * time.Second)
	//}
}