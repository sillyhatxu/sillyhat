package main

import (
	"fmt"
	"sillyhat/basic/stacker/stack"
)

func main() {
	var testStack stack.Stack
	testCap := testStack.Cap()
	fmt.Println("stack cap is ",testCap)
	for i:=0;i < 16;i++ {
		fmt.Println("stack cap is ",testStack.Cap(),";length is ",testStack.Len())
		testStack.Push("i + 1 = " + fmt.Sprintf("%d", i + 1))
	}
	fmt.Println("stack cap is ",testStack.Cap(),";length is ",testStack.Len())
	for i:=0;i < 16;i++ {
		testStack.Push("i + 1 = " + fmt.Sprintf("%d", i + 100))
	}
	fmt.Println("stack cap is ",testStack.Cap(),";length is ",testStack.Len())
}