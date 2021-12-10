package main

import (
	"fmt"
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

//messy, needs cleaning :)
func p2(){
	ireduced := input
	iireduced := input
	runNor, runInv := true, true
	for i := 0; i < len(input[0]); i++ {

		cache := []string{}
		cacheb := []string{}

		if(runNor){
		needle := columnBit(i, ireduced, false)

		for _, bits := range ireduced {
			if bits[i] == byte(needle) {
				cache = append(cache, bits)
			}
			ireduced = cache
		}
		}

		if(runInv){
		needle2 := columnBit(i, iireduced, true)
			for _, bit := range iireduced {
				if(bit[i] == byte(needle2)) {
					cacheb = append(cacheb, bit)
				}
			}
			iireduced = cacheb
		}

		if len(ireduced) == 1 {
			runNor = false
		}

		if len(iireduced) == 1{
			runInv = false
		}
	}
	fmt.Printf("\n %s %d %s","PART2 ANSWER : ", (u.StrBinaryToDecimal(ireduced[0]) * u.StrBinaryToDecimal(iireduced[0]))," <- ")
}

func columnBit(p int, input []string, inv bool) rune {
	io := 0
	ii := 0

	for _, s := range input {
		if s[p] == '1'{
			ii++
		}else {
			io++
		}
		println("ii", ii)
	}

	if io <= ii {
		if(inv){
			return '0'
		}
		return '1'
	}
	if(inv){
		return '1'
	}
	return '0'
}
