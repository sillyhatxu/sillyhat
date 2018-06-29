package main

import (
	"log"
	"strconv"
)

type IDuck struct {
	id string
	iDuck interface{}
}

type IDuckArray interface {
	getIDuckArray() []IDuck
}

func DuckDance(array IDuckArray) {
	for _,iDuck := range array.getIDuckArray(){
		log.Println(iDuck.id)
	}
}

type Bird struct {
	ProductId int
	ProductName string
}

func (b *Bird) Id() string {
	return strconv.Itoa(b.ProductId)
}

type BirdArray struct {
	IDuckArray
	birdArray []Bird

}

func (array *BirdArray) getIDuckArray() []IDuck {
	//var iduckArray []IDuck
	//for _,b := range array.birdArray{
	//	iduckArray = append(iduckArray,b)
	//}
	return nil
}


func main() {
	var array BirdArray
	var birdArray []Bird
	for i:=0;i<10;i++{
		birdArray = append(birdArray, *&Bird{ProductId:i,ProductName:"Test-" + strconv.Itoa(i)})
	}
	//for _,b := range birdArray{
	//	duckArray = append(duckArray,b)
	//}
	array.birdArray = birdArray
	DuckDance(array)
}