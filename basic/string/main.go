package main

import (
	"strconv"
	"fmt"
	"strings"
)

func check(error error){
	if error!= nil{
		panic(error)
	}
}



func main() {
	fmt.Println(strings.Replace("/oink /oink oi /nk / ", " ", "ky", -1))
	fmt.Println(strings.Replace("  /oink /oink oi /nk / ", " ", "_",-1))
	fmt.Println(strings.Replace("  /oink /oink oi /nk / ", "/", "_",-1))
	fmt.Println(strings.Replace(strings.Replace("  /oink /oink oi /nk / ", "/", "_",-1), " ", "_",-1))

	//string到int
	int,err:=strconv.Atoi("5")
	check(err)
	//string到int64
	int64, err := strconv.ParseInt("856", 10, 64)
	//int到string
	string:=strconv.Itoa(int)
	//int64到string
	test:=strconv.FormatInt(int64,10)
	fmt.Println(test)
	//string到float32(float64)
	float,err := strconv.ParseFloat(string,32/64)
	fmt.Println(float)
	//float到string
	//var float32Test float32
	var float64Test float64
	//float32Test = '5'
	float64Test = 2022
	//test32 := strconv.FormatFloat(float32Test, 'E', -1, 32)
	test64 := strconv.FormatFloat(float64Test, 'f', -1, 64)
	fmt.Println(test64)
	fmt.Println("----------------")
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	// 1.00123e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	// 1.00123E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	// 100.12346
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'b', 30, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 30, 32))
	// 1.001234588623046875000000000000e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 30, 32))
	// 1.001234588623046875000000000000E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 30, 32))
	// 100.123458862304687500000000000000
	fmt.Println(strconv.FormatFloat(f, 'g', 30, 32))
	// 100.1234588623046875
	fmt.Println(strconv.FormatFloat(f, 'G', 30, 32))
	// 100.1234588623046875
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)

	fmt.Println("----------------")
	var testToString float64
	testToString = 2025
	fmt.Println(strconv.FormatFloat(testToString, 'f', 0, 64))
}

