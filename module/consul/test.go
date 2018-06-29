package main


import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
	"log"
)

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}
//
func registerServer() {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	checkPort := 8301
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "serverNode_1"
	registration.Name = "serverNode"
	registration.Port = 9527
	registration.Tags = []string{"serverNode"}
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

	http.HandleFunc("/check", consulCheck)
	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)

}

func main() {
	registerServer()
	// http://127.0.0.1:18001/health
	//client, _ := consulapi.NewClient(consulapi.DefaultConfig())
	//kv := client.KV()
	//// PUT a new KV pair
	//p := &consulapi.KVPair{Key: "foo", Value: []byte("test")}
	//_, err := kv.Put(p, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Lookup the pair
	//pair, _, err := kv.Get("foo", nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("KV: %v", pair)
	//config := consulapi.DefaultConfig()
	//config.Address = "127.0.0.1:8080"
	////config.Address = "127.0.0.1:18001"
	//config.HttpAuth = &consulapi.HttpBasicAuth{Username: "guest", Password: "secret"}
	//client, err := consulapi.NewClient(config)
	//if err != nil {
	//	panic(err)
	//}
	//kv := client.KV()
	////d := &consulapi.KVPair{Key: "sites/1/domain", Value: []byte("example.com")}
	////kv.Acquire(d, nil)
	//kvp, qm, err := kv.Get("health", nil)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(kvp.Value))
	//}
	//fmt.Println(qm)
	//
	////pair, _, err := kv.Get("/health", nil)
	////fmt.Printf("KV: %v\n", pair)
	//log.Println("-----------------")

	//checkPort := 8080
	//http.HandleFunc("/check", consulCheck)
	//http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)

}