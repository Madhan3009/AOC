package utils


import(
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
	"path/filepath"
)
func Read(fileName string)([]int,[]int,error){
	path := buildPath(fileName)
	input1,input2,err := readShit(path)
	return input1,input2,err
}
func buildPath(fileName string) string{
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	
	fullPath := filepath.Join(cwd, fileName)

	return fullPath
}
func readShit(path string)([]int,[]int,error){
	data,err := os.ReadFile(path)
	if(err!=nil){
		log.Fatal(err)
	}


	
	lines := strings.Split(strings.TrimSpace(string(data)),"\n")
	left := make([]int,0,len(lines))
	right := make([]int,0,len(lines))
	
	for _,line := range lines{
		fields := strings.Fields(line)
		
		lString := fields[0]
		rString := fields[1]
		
		val1,_ := strconv.Atoi(lString)
		left = append(left,val1)

		val2,_ := strconv.Atoi(rString)
		right = append(right,val2)
	
	}

	return left,right,nil
}


