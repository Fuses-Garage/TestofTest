package TestofTest

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Identifer string
	args      []int
	want      int
}

func TestMax(t *testing.T) {
	tc := []TestCase{{"test01", []int{8, 4}, 12}, {"test02", []int{2, 3}, 5}, {"test03", []int{1, 3, 5}, 9}, {"test04", []int{1, 2, 3, 4}, 10}}
	for _, v := range tc {
		t.Run(v.Identifer, v.Testing)
	}
}
func (c TestCase) Testing(t *testing.T) {
	out := Add(c.args)
	successed := out == c.want
	if successed {
		fmt.Printf("Success want:%d out:%d\n", c.want, out)
	} else {
		t.Errorf("Failed want:%d out:%d\n", c.want, out)
	}
}
