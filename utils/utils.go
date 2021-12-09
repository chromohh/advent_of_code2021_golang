package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//parse file to strarr
func ParseFileToStringArr(path string)([]string, error){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("epic fail", err)
		return nil, err
	}
	defer file.Close()
	var lines []string
	buffer := bufio.NewScanner(file)
	for buffer.Scan(){
		lines = append(lines, buffer.Text())
	}

	log.Print("file parsed")
	return lines, buffer.Err()
}

//parse file to int arr
func ParseFileToIntArr(path string)([]int, error){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("fail", err)
		return nil, err
	}
	defer file.Close()
	var lines []int
	buffer := bufio.NewScanner(file)
	for buffer.Scan(){
		parsedInt, err := strconv.Atoi(buffer.Text());
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		lines = append(lines, parsedInt)
	}
	return lines, err
}

func ParseFileToByteBuff(path string)([]byte, error){
	return os.ReadFile(path);
}

func StringToInt(s string)(int){
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("CANT PARSE THIS STRING :D")
	}
	return i;
}

func IntToString(s int)(string){
	i := strconv.Itoa(s)
	return i;
}

func StrBinaryToDecimal(i string)(int){
	output, err := strconv.ParseInt(i, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(output)
}

func StrToIntArr(s string) []int {
    strs := strings.Split(s, "")
    res := make([]int, len(strs))
    for i := range res {
        res[i], _ = strconv.Atoi(strs[i])
    }
    return res
}

func NumArrToString(s []int)(string){
	rs := ""
	for _, v := range s{
		rs += strconv.Itoa(v)
	}
	return rs;
}

func RemoveRegexMatchStrArr(slice []string, pattern regexp.Regexp) ([]string) {
	offset := len(slice) - 1;
	for i := 0; i <= offset; i++ {
		if pattern.MatchString(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			offset--
		}
	}
	return slice
}
