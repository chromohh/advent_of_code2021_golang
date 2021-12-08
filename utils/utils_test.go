package utils

import (
	"reflect"
	"testing"
)

const intfile = "./testinputs/ints.txt"
func TestIntParser(t *testing.T) {
	var expected = []int{123, 123, 4, 9191}
	actual, err := ParseFileToIntArr(intfile);

	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(actual, expected){
		t.Errorf("integer array does not match")
	}
}

const stringfile = "./testinputs/ints.txt"

func TestStringPaser(t *testing. T){
	var expected = []string{"sdsdkask", "dksakd", "las", "dlacla", "clsad", "la", "lacs", "lda"}
	actual, err := ParseFileToStringArr(stringfile)

	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(actual, expected){
		t.Errorf("string array does not match");
	}
}
