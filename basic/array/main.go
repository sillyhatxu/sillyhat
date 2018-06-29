package main

import "log"

func testArray()  {
	grid1 := make([][]int,10)
	for i := range grid1{
		grid1[i] = make([]int,3)
	}
	grid1[1][0],grid1[1][1],grid1[1][2] = 8,2,6
	log.Println(grid1)

	grid2 := make([]int,10)
	for i := range grid1{
		grid2[i] = make([]string,3)
	}
	grid2[1][0],grid2[1][1],grid2[1][2] = 8,2,6
	log.Println(grid1)
}


func main() {
	testArray()
}
