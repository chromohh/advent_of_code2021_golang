package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	file, err := os.Create(path)
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
