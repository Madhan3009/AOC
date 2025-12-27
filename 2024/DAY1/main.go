package main

import(
	"fmt"
	"AOC2024/utils"
	"log"
	"slices"
)
func main(){
	input1,input2,err := utils.Read("message.txt")
	if err != nil{
		log.Fatal("Error at Reading file")
	}
	slices.Sort(input1)
	slices.Sort(input2)

	distance := 0
	for i,_ := range input1{
		sum := input1[i] - input2[i]
		if(sum<0){
			sum = -sum
		}
		distance += sum
	}

	fmt.Print(distance)
}

