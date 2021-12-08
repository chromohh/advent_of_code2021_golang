package utils

import (
	"reflect"
	"testing"
	"fmt"
)

const intfile = "./testinputs/ints.txt"
func TestIntParser(t *testing.T) {
	var expected = []int{123, 123, 9191, 12345, 1, 2, 3}
	actual, err := ParseFileToIntArr(intfile);

	if err != nil {
		t.Fatal("error = nul")
	}
	if !reflect.DeepEqual(actual, expected){
		t.Errorf(err.Error());
		fmt.Print(err)
		t.Errorf("integer array does not match")
	}
}

const stringfile = "./testinputs/strings.txt"
func TestStringPaser(t *testing.T){
	var expected = []string{"sdsdkask", "dksakd", "las", "dlacla"}
	actual, err := ParseFileToStringArr(stringfile)

	if err != nil {
		t.Fatal("err = nil")
	}
	if !reflect.DeepEqual(actual, expected){
		t.Errorf("string array does not match");
	}
}
