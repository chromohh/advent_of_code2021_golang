package main

import (
	"fmt"
	"regexp"

	"chromo.com/aocgolang/utils"
	u "chromo.com/aocgolang/utils"
)

const PATH = "../day3/input.txt"
var input, err = u.ParseFileToStringArr(PATH)

func main() {
	bits := []int{}
	for _, str := range input{
		intsplice := utils.StrToIntArr(str)
		for i, n := range intsplice{
			if(len(bits) < i+1){
				bits = append(bits, n)
			}else{
				bits[i] += n;
			}
		}
	}
	total := len(input);
	final1 := "";
	final2 := "";
	for _, bit := range bits{
		if bit > (total/2){
			final1 += u.IntToString(1)
			final2 += u.IntToString(0)
		}else{
			final1 += u.IntToString(0)
			final2 += u.IntToString(1)
		}
	}

	result := u.StrBinaryToDecimal(final1) * u.StrBinaryToDecimal(final2)
	fmt.Printf("\n %s %d %s","PART1 ANSWER: ", result, " <- ")
    p2();
}

func p2(){
	//copy input
	//loop amnt of bits
	//loop new input copy
	//get columnBit
	//if columnbit = x = append(type, elems)
	//

}

func columnBit(pos int, input []string){
	//string[index] ?
	//otherwise byte buff
}
