package main

import "fmt"

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
	for i:= 1;i < 5000 ; i++ {
		fmt.Println(test(i))
	}
}
