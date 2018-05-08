package main

import (
	"flag"
	"fmt"
)

func main() {
	var enviroment string
	flag.StringVar(&enviroment, "enviroment", "dt", "enviroment")
	flag.Parse()
	fmt.Printf("enviroment : %v \n",enviroment)
}



