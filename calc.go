package TestofTest

func Add(x []int) int { //intスライスの合計を返す
	sum := 1              //合計を格納する変数
	for _, v := range x { //スライスの中身を1つずつ取り出す
		sum += v //加える
	}
	return sum //合計を返す
}
