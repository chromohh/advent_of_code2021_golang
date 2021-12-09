package main

import(
	u "chromo.com/aocgolang/utils"
    "fmt"
	"regexp"
	"strconv"
)

const PATH = "../day2/input.txt"
var input, _ = u.ParseFileToStringArr(PATH)

func main() {
	cordIX, cordY := 0, 0
	r, _ := regexp.Compile("[0-9]")

	for _, element := range input{
		s := r.FindString(element)
		n, _ := strconv.Atoi(s);
			switch element[:1]{
		    case "u": cordIX -= n
		    case "f": cordY += n
			case "d": cordIX += n
		}
	}
	fmt.Printf("\n %s %d", "PART1: Result " , (cordIX*cordY))
}
