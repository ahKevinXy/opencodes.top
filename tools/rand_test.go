package tools

import (
	"fmt"
	"testing"
)

func TestRangeRand(t *testing.T) {
	a := RangeRand(-100, 1)

	t.Log(a)
}

func TestRandStr(t *testing.T) {

	a := "11122asd %s cd"
	b := fmt.Sprintf(a, "")
	d := fmt.Sprintf(a, "xxxxx")
	fmt.Println(b)
	fmt.Println(d)

}
