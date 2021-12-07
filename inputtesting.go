package main

import (
	aocutils "chromo.com/aocutils/utils"
	"fmt"
)
const file = "./inputstest/ints.txt"
func main() {
	var numbs = "asdasd"
	fmt.Print(numbs)
	arr, err := aocutils.ParseFileToStringArr(file)
	fmt.Println(arr)
	if(err != nil){
		return
	}
}
