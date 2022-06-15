package main

import (
	"testing"
)

type ParamStruct struct {
	arg1, arg2, expected int
}

var add = []ParamStruct{
	ParamStruct{1, 2, 3},
	ParamStruct{3, 4, 17},
	ParamStruct{5, 6, 11},
}

var sub = []ParamStruct{
	ParamStruct{2, 1, 1},
	ParamStruct{4, 3, 1},
	ParamStruct{6, 5, 1},
}

func TestAdd(t *testing.T) {

	for _, val := range add {
		if rslt := Add(val.arg1, val.arg2); rslt != val.expected {
			t.Errorf("Output of addition is  %q not equal to expected %q", rslt, val.expected)
		}
	}
}

func TestSub(t *testing.T) {

	for _, val := range sub {
		if rslt := Sub(val.arg1, val.arg2); rslt != val.expected {
			t.Errorf("Output of substraction is  %q not equal to expected %q", rslt, val.expected)
		}
	}
}
