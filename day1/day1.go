package main

import(
	u "chromo.com/aocgolang/utils"
	"fmt"
)

const PATH = "../day1/input.txt"
var input, err = u.ParseFileToIntArr(PATH)
func main() {
	if err != nil {
		fmt.Print("NULL")
	}
	var ups = 0;

	for index, element := range input{
		if(index != 0){
			if(element < input[index - 1]){
				ups += 1
			}
		}
	}

	fmt.Printf("\n %s %d %s","PART1: Depth increased: ", ups, " times")
    p2();
}

func p2(){
	var wups = 0;
	for index, element := range input{
		if(index <= (len(input) - 4)){
			w1 := element + input[index + 1] + input[index + 2];
			w2 := input[index + 1] + input[index + 2] + input[index + 3];
			if(w2 > w1){
				wups += 1
			}
		}
	}
	fmt.Printf("\n %s %d", "PART2: Windowed depth increased: ", wups);
}
