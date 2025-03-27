package parser

import (
	"fmt"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type Base struct {
		A int `json:"a"`
	}

	type p struct {
		Base

		B int `json:"b"`
	}

	d := p{
		Base: Base{
			A: 1,
		},
		B: 2,
	}

	result := StructToMap(d)
	fmt.Printf("%T \n", result)

	if result["a"] != 1 {
		t.Error("a should be 1")
	}

	if result["b"] != 2 {
		t.Error("b should be 2")
	}
}
