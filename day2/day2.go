package main

import (
	"fmt"
	"regexp"
	u "chromo.com/aocgolang/utils"
)

const PATH = "../day2/input.txt"
var input, _ = u.ParseFileToStringArr(PATH)
var r,_ = regexp.Compile("[0-9]")

func main() {
	cordIX, cordY := 0, 0
	for _, element := range input{
		s := r.FindString(element)
		n := u.StringToInt(s)
			switch element[:1]{
		    case "u": cordIX -= n
		    case "f": cordY += n
			case "d": cordIX += n
		}
	}
	fmt.Printf("\n %s %d", "PART1: Result " , (cordIX*cordY))
	p2()
}

func p2(){
	a, h, d := 0, 0, 0
	for _, element := range input{
		s := r.FindString(element)
		n := u.StringToInt(s)
			switch element[:1]{
				case "u":
				a -= n
				case "f":
				h += n
				d += (a*n)
				case "d":
				a += n
		}
	}
	fmt.Printf("\n %s %d", "PART2 : Resulet ", (h * d))
}
