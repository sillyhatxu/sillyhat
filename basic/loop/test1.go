package main

import "fmt"

func testBreak() {
Loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break Loop
		}
		fmt.Println(i)
	}
}

func testContinue() {
Loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue Loop
		}
		fmt.Println(i)
	}
}

func testGoto() {
Loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto Loop
		}
		fmt.Println(i)
	}
}

func main() {
	testBreak()
	testContinue()
	testGoto()
}
