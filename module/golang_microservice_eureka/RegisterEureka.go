package main
//
//import (
//	"sync"
//	"os/signal"
//	"os"
//	"syscall"
//	"log"
//	"net/http"
//	"github.com/eriklupander/eeureka"
//)
//
//
//func handleSigterm() {
//	c := make(chan os.Signal, 1)          // Create a channel accepting os.Signal
//	// Bind a given os.Signal to the channel we just created
//	signal.Notify(c, os.Interrupt)        // Register os.Interrupt
//	signal.Notify(c, syscall.SIGTERM)     // Register syscall.SIGTERM
//
//	go func() {                           // Start an anonymous func running in a goroutine
//		<-c                           // that will block until a message is recieved on
//		eureka.Deregister()           // the channel. When that happens, perform Eureka
//		os.Exit(1)                    // deregistration and exit program.
//	}()
//}
//
//func startWebServer(port string) {
//	router := service.NewRouter()
//	log.Println("Starting HTTP service at " + port)
//	err := http.ListenAndServe(":" + port, router)
//	if err != nil {
//		log.Println("An error occured starting HTTP listener at port " + port + ": " + err.Error())
//	}
//}
//
//
//func main() {
//	handleSigterm()                              // Handle graceful shutdown on Ctrl+C or kill
//
//	go startWebServer("18010")                          // Starts HTTP service  (async)
//
//	eureka.Register()                            // Performs Eureka registration
//
//	go eureka.StartHeartbeat()                   // Performs Eureka heartbeating (async)
//
//	// Block...
//	wg := sync.WaitGroup{}                       // Use a WaitGroup to block main() exit
//	wg.Add(1)
//	wg.Wait()
//}