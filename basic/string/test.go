package main

import (
	"fmt"
	"strings"
)

func test(a int) string {
	//fmt.Println(a)
	//前置补0
	//fmt.Printf("%02d\n", a)
	//test,_ := fmt.Printf("%02d", a)
	//fmt.Println(test)
	//fmt.Printf("%0*d", 3, a)
	return fmt.Sprintf("%02d",a)
}
func main() {
	//for i:= 1;i < 5000 ; i++ {
	//	fmt.Println(test(i))
	//}
	fmt.Println(strings.Replace("/oink /oink oi /nk / ", " ", "ky", 2))
	fmt.Println(strings.Replace("  /oink /oink oi /nk / ", " ", "_",-1))
	fmt.Println(strings.Replace("  /oink /oink oi /nk / ", " ", "_",0))
	fmt.Println(strings.Replace(`{"sizes":["One Size"],"measurements":["PTP","Waist","Length"],"values":[["21#INCH#"],["21#INCH#"],["20#INCH#"]]}`, "#INCH#", "\"",-1))
}
