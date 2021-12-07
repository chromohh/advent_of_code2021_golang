package main

import (
	"fmt"
	"chromo.com/aocutils/utils"
)
const file = "/day1.txt"
func main() {
	var numbs = "asdasd"
	fmt.Print(numbs)
	arr, err := utils.ParseFileToStringArr(file)
	fmt.Println(arr)
	if(err != nil){
		return
	}
}
