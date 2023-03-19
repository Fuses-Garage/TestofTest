package TestofTest

import (
	"fmt"
	"testing"
)

type TestCase struct { //テストケース用構造体
	identifier string //名前
	args       []int  //入力値
	want       int    //想定した出力
}

func TestAdd(t *testing.T) {
	tc := []TestCase{
		{"test01", []int{8, 4}, 12},
		{"test02", []int{2, 3}, 5},
		{"test03", []int{1, 3, 5}, 9},
		{"test04", []int{1, 2, 3, 4}, 10},
	}
	for _, v := range tc {
		t.Run(v.identifier, v.Testing)
	}
}

func (c TestCase) Testing(t *testing.T) { //テストを行う関数
	out := Add(c.args)       //関数の出力値を取得
	success := out == c.want //テストは成功か？
	if success {             //成功なら
		fmt.Printf("Success want:%d out:%d\n", c.want, out) //ログを出す
	} else { //失敗したら
		t.Errorf("Failed want:%d out:%d\n", c.want, out) //失敗のログを出す
	}
}
