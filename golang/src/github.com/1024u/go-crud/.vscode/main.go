//学べること...
// for range...変数indexにはループ回数が、変数valueには配列の値が、変数userNamesの中から順番に代入される。
//make()...スライス、マップ、およびチャネルを生成（初期化）するのに使用される。
// 		　　スライスの場合、事前にスライスのcap（容量）を第三引数にて指定できる。
//　　　　　第3引数で指定した数字分、appendで値を追加できる。
//　　　　　事前に、どれくらい値を入れるか分かっている場合はmake()を使うと容量を食わない

package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//make()...事前にスライスのcap（容量）を第三引数にて指定できる。
	//　　　　　第3引数で指定した数字分、appendで値を追加できる。
	//　　　　　事前に、どれくらい値を入れるか分かっている場合はmake()を使うと容量を食わない
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//make()を使ったmapの作成
	//map[string]float...空のキーとそれに紐づく値を作成している。map[キー]紐づく値
	//courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	//キーと、それに紐づく値を変数courseRatingsに追加している。
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	//floatMap型に紐づくoutput()メソッドを呼び出している。
	courseRatings.output()

	//for range...変数indexにはループ回数が、変数valueには配列の値が、
	// 			　変数userNamesの中から順番に代入される。
	for index, value := range userNames {
		// ...
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

}
