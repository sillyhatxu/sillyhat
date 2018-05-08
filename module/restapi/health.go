package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

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
	router := mux.NewRouter()
	router.HandleFunc("/health", health).Methods("GET")
	log.Fatal(http.ListenAndServe(":18001", router))
}


func main() {
	startHealthApi()
}