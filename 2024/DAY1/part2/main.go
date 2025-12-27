package main

import(
	"fmt"
	"AOC2024/utils"
	"log"
)
func main(){

	input1,input2,err := utils.Read("message.txt")
	if err != nil{
		log.Fatal("Error at Reading file")
	}
	mapper := make(map[int]int)	
	
	for _,v := range input2{
		mapper[v]++
	}

	distance := 0

	for _,v := range input1{
			value := mapper[v]
			product := v * value
			distance += product
	}

	fmt.Print(distance)
}

