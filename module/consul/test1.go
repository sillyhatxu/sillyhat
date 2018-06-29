package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	//client, _ := consuldiscovery.NewClient(consuldiscovery.DefaultConfig())
	//
	//services, _ := client.CatalogServices()
	//fmt.Printf("Services: %#v\n\n", services)
	//
	//for _, service := range services {
	//	serviceNodes, _ := client.CatalogServiceByName(service.Name)
	//	fmt.Printf("%s: %#v\n", service.Name, serviceNodes)
	//}

	// Get a new client, with KV endpoints
	client, _ := consulapi.NewClient(consulapi.DefaultConfig())
	kv := client.KV()
	// PUT a new KV pair
	p := &consulapi.KVPair{Key: "foo", Value: []byte("test")}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v", pair)
	// http://127.0.0.1:18001/health
}